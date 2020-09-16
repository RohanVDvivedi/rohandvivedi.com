package chatter

import (
	"golang.org/x/net/websocket"
	"time"
	"errors"
)

type ChatUser struct {
	Id
	Name
	PublicKey string

	MessagesPendingToBeSent *ChatMessageQueue

	CConn *ChatConnection

	ChatGroups map[string]*ChatGroup
}

func NewChatUser(name string, publicKey string) *ChatUser {
	user := &ChatUser{
		Id: Id{GetNewChatUserId()},
		Name: Name{name},
		PublicKey: publicKey,
		MessagesPendingToBeSent: NewChatMessageQueue(),
		ChatGroups: make(map[string]*ChatGroup),
	}
	return user
}

func (user *ChatUser) SetChatConnection(CConn *ChatConnection) {
	user.CConn = CConn
	go user.LoopOverChannelToPassMessages()
}

func (user *ChatUser) SendMessage(msg ChatMessage) {
	if(msg.To == user.GetId()) {
		if(user.CConn != nil) {
			user.CConn.SendMessage(msg)
		} else {
			user.MessagesPendingToBeSent.Push(msg)
		}
	}
}

func (user *ChatUser) ReceiveMessage() (ChatMessage, error) {
	msg, err := user.CConn.ReceiveMessage()
	if(err != nil) {	// this could mean, connection closed or malformed chatMessage packet
		return msg, err;
	} else if(msg.From != user.GetId()) { // user maliciously pretending to be someone else
		return ChatMessage{}, errors.New("Malicious user: wrong user name in from attribute")
	}
	return msg, nil
}

func (user *ChatUser) LoopOverChannelToPassMessages() {
	for (true) {
		msg := cconn.MessagesToBeSent.Top()
		if(msg.To == user.Name) {
			err := ChatMessageCodec.Send(user.Connection, msg)
			if(err != nil) { // this could mean, connection closed or lost
				user.Destroy()
				return
			}
			user.LastMessage = time.Now()
		}
		cconn.MessagesToBeSent.Pop()
	}
}

func (user *ChatUser) Destroy() {
	close(user.MessagesToBeSent)
	user.Connection.Close()
}