package chatter

import (
	"golang.org/x/net/websocket"
	"time"
	"errors"
)

type ChatUser struct {
	ChatterBoxIndentity
	PublicKey string

	MessagesToBeSent chan ChatMessage
	MessagesReceived chan ChatMessage

	CConn *ChatConnection

	ChatGroups map[string]*ChatGroup
}

func NewChatUser(name string, publicKey string) *ChatUser {
	user := &ChatUser{
		ChatterBoxIndentity: ChatterBoxIndentity{Id: GetNewChatUserId(), Name: name},
		PublicKey:publicKey,

		MessagesToBeSent: make(chan ChatMessage, 10),
		MessagesReceived: make(chan ChatMessage, 10),

		ChatGroups:make(map[string]*ChatGroup),
	}
	go user.LoopOverChannelToPassMessages()
	return user
}

func (user *ChatUser) SendMessage(msg ChatMessage) {
	if(msg.To == user.GetId()) {
		user.CConn.SendMessage(msg)
	}
}

func (user *ChatUser) ReceiveMessage() (ChatMessage, error) {
	msg <- user.CConn.MessagesReceived
	err := ChatMessageCodec.Receive(user.Connection, &msg)
	if(err != nil) {	// this could mean, connection closed or malformed chatMessage packet
		return msg, err;
	}
	user.LastMessage = time.Now()
	if(msg.From != user.Name) { // user maliciously pretending to be someone else
		return ChatMessage{}, errors.New("Malicious user: wrong user name in from attribute")
	}
	return msg, nil
}

func (user *ChatUser) LoopOverChannelToPassMessages() {
	for msg := range user.MessagesToBeSent {
		if(msg.To == user.Name) {
			err := ChatMessageCodec.Send(user.Connection, msg)
			if(err != nil) { // this could mean, connection closed or lost
				user.Destroy()
				return
			}
			user.LastMessage = time.Now()
		}
	}
}

func (user *ChatUser) Destroy() {
	close(user.MessagesToBeSent)
	user.Connection.Close()
}