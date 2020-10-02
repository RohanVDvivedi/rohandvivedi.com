package chatter

import (
	"golang.org/x/net/websocket"
	"errors"
	"rohandvivedi.com/src/session"
)

type ChatConnection struct {
	Id

	Connection *websocket.Conn
	
	// this is the user that this connection belongs to
	User *ChatUser
}

func NewChatConnection(Connection *websocket.Conn) *ChatConnection {
	cconn := &ChatConnection {
		Id: Id{GetNewChatConnectionId()},
		Connection: Connection,
	}

	// add ChatConnection to the session values
	session.GlobalSessionStore.GetExistingSession(cconn.Connection.Request()).
	ExecuteOnValues(func (values map[string]interface{}, add interface{}) interface{} {
		chatConnsIntr, found := values["chat_conns"]
		chatConns, isValid := chatConnsIntr.(map[string]*ChatConnection)
		if(!found || !isValid) {
			chatConns = map[string]*ChatConnection{}
		}
		chatConns[cconn.GetId()] = cconn
		values["chat_conns"] = chatConns
		return nil
	}, nil)

	return cconn
}

func (cconn *ChatConnection) GetDetailsAsString() string {
	return cconn.GetId()
}

func (cconn *ChatConnection) SendMessage(msg ChatMessage) error {
	msg.OriginConnection = ""
	return ChatMessageCodec.Send(cconn.Connection, msg)
}

func (cconn *ChatConnection) ReceiveMessage() (ChatMessage, error) {
	msg := ChatMessage{}
	err := ChatMessageCodec.Receive(cconn.Connection, &msg)
	if(err != nil) {
		// this could mean, connection closed or malformed chatMessage packet, both of which are fatal
		return msg, err
	}
	msg.OriginConnection = cconn.GetId()

	// the message sender field must be empty or user id of the user, if he/she is logged in, No fraud must be allowed
	msgSenderFromFieldValid := (msg.From == cconn.GetId()) || (cconn.User != nil && msg.From == cconn.User.GetId())
	if (!msgSenderFromFieldValid){
		return ChatMessage{}, errors.New("ERROR user attempting identity theft")
	}

	// # TODO this needs to be moved to users logic, because if a message is sent to or not is a users responsibility to acknoledge
	// the receive function is responsible to send the SENT receipt for every message, that is received and is processable
	if(IsChatId(msg.To)) {
		ChatMessageCodec.Send(cconn.Connection, ChatMessage{From: msg.To, To: cconn.GetId(), Message: msg.MessageId + "-SENT", ContextId: msg.MessageId})
	}

	return msg, nil
}

func (cconn *ChatConnection) Destroy() {
	// remove chat connection from session store
	session.GlobalSessionStore.GetExistingSession(cconn.Connection.Request()).
	ExecuteOnValues(func (values map[string]interface{}, add interface{}) interface{} {
		chatConnsIntr, found := values["chat_conns"]
		chatConns, isValid := chatConnsIntr.(map[string]*ChatConnection)
		if(!found || !isValid) {
			chatConns = map[string]*ChatConnection{}
		}
		delete(chatConns, cconn.GetId())
		values["chat_conns"] = chatConns
		return nil
	}, nil)

	if(cconn.User != nil) {
		BreakConnectionFromUser(cconn, cconn.User)
	}

	cconn.Connection.Close()
}

/* Joinery methods */
func (cconn *ChatConnection) HasChatUser() bool {
	return cconn.User != nil
}
func (cconn *ChatConnection) GetChatUser() *ChatUser {
	return cconn.User
}
func (cconn *ChatConnection) SetChatUser(c *ChatUser) {
	cconn.User = c
}
func (cconn *ChatConnection) RemoveChatUser() {
	cconn.User = nil
}


/* Below methods update modify session values and must be called only 
while the corresponding socket connection is active  atleast as per the chat manager */

func (cconn *ChatConnection) GetNameAndPublicKey() (string, string, bool) {
	return GetNameAndPublicKeyFromSession(session.GlobalSessionStore.GetExistingSession(cconn.Connection.Request()))
}

func (cconn *ChatConnection) SetNameAndPublicKey(name string, publicKey string) {
	InsertNameAndPublicKeyToSession(session.GlobalSessionStore.GetExistingSession(cconn.Connection.Request()), name, publicKey)
}

func (cconn *ChatConnection) RemoveNameAndPublicKey() {
	RemoveNameAndPublicKeyFromSession(session.GlobalSessionStore.GetExistingSession(cconn.Connection.Request()))
}