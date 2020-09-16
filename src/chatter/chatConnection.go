package chatter

import (
	"golang.org/x/net/websocket"
	"errors"
)

type ChatConnection struct {
	Id

	Connection *websocket.Conn
	
	// this is the user that this connection belongs to
	User *ChatUser
}

func NewChatConnection(Connection *websocket.Conn, User *ChatUser) *ChatConnection {
	return &ChatConnection {
		Id: Id{GetNewChatConnectionId()},
		Connection: Connection,
		User: User,
	}
}

func (cconn *ChatConnection) SendMessage(msg ChatMessage) error {
	if(msg.To == cconn.GetId() || msg.To == cconn.User.GetId()) {
		return ChatMessageCodec.Send(cconn.Connection, msg)
	}
	return errors.New("ERROR in routing")
}

func (cconn *ChatConnection) ReceiveMessage() (ChatMessage, error) {
	msg := ChatMessage{}
	err := ChatMessageCodec.Receive(cconn.Connection, &msg)
	if(err != nil) {	// this could mean, connection closed or malformed chatMessage packet
		return msg, err
	} else if (msg.From != cconn.GetId() && msg.From != cconn.User.GetId()) {
		return ChatMessage{}, errors.New("ERROR user attempting identity theft")
	}
	return msg, nil
}

func (cconn *ChatConnection) Destroy() {
	cconn.Connection.Close()
}