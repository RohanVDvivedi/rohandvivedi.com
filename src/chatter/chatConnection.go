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

func NewChatConnection(Connection *websocket.Conn) *ChatConnection {
	return &ChatConnection {
		Id: Id{GetNewChatConnectionId()},
		Connection: Connection,
	}
}

func (cconn *ChatConnection) SendMessage(msg ChatMessage) error {
	// if this connection was responsible for generating the message, then do not send it again
	if(msg.OriginConnection != cconn.GetId()) {
		return ChatMessageCodec.Send(cconn.Connection, msg)
	}
	return nil
}

func (cconn *ChatConnection) ReceiveMessage() (ChatMessage, error) {
	msg := ChatMessage{}
	err := ChatMessageCodec.Receive(cconn.Connection, &msg)
	if(err != nil) {	// this could mean, connection closed or malformed chatMessage packet
		return msg, err
	}

	// the message sender field must be empty or user id of the user
	msgSenderFromFieldValid := (msg.From == cconn.GetId()) || (cconn.User != nil && msg.From == cconn.User.GetId())
	if (!msgSenderFromFieldValid){
		return ChatMessage{}, errors.New("ERROR user attempting identity theft")
	}

	msg.OriginConnection = cconn.GetId()
	return msg, nil
}

func (cconn *ChatConnection) Destroy() {
	cconn.Connection.Close()
}