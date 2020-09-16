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

	ChatConnections map[string]*ChatConnection

	ChatGroups map[string]*ChatGroup
}

func NewChatUser(name string, publicKey string) *ChatUser {
	user := &ChatUser{
		Id: Id{GetNewChatUserId()},
		Name: Name{name},
		PublicKey: publicKey,
		MessagesToBeSent: NewChatMessageQueue(),
		ChatGroups: make(map[string]*ChatGroup),
	}
	return user
}

func (user *ChatUser) SendMessage(msg ChatMessage) {
	if(msg.To == user.GetId()) {
		user.MessagesToBeSent.Push(msg)
	}
}

func (user *ChatUser) LoopToForwardMessages() {
	for (true) {
		msg := cconn.MessagesToBeSent.Top()
		if(msg.To == user.GetId()) {
			for _, conn := range grp.ChatConnections {
				conn.SendMessage(msg)
			}
		}
		cconn.MessagesToBeSent.Pop()
	}
}

func (user *ChatUser) Destroy() {
	close(user.MessagesToBeSent)
	user.Connection.Close()
}