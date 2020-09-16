package chatter

import (
	"time"
)

type ChatGroup struct {
	Id
	Name

	MessagesToBeSent chan ChatMessage

	ChatUsers map[string]*ChatUser

	LastMessage time.Time
}

func NewChatGroup(name string) *ChatGroup {
	grp := &ChatGroup{
		Id: Id{GetNewChatUserId()},
		Name: Name{name},
		MessagesToBeSent: make(chan ChatMessage, 10),
		ChatUsers: make(map[string]*ChatUser),
		LastMessage: time.Now(),
	}
	go grp.LoopOverChannelToPassMessages()
	return grp
}

func (grp *ChatGroup) SendMessage(msg ChatMessage) {
	grp.MessagesToBeSent <- msg
}

func (grp *ChatGroup) LoopOverChannelToPassMessages() {
	for msg := range grp.MessagesToBeSent {
		if(msg.To == grp.Id) {
			for _, user := range grp.ChatUsers {
				user.SendMessage(msg)
			}
			grp.LastMessage = time.Now()
		}
	}
}

func (grp *ChatGroup) Destroy() {
	close(grp.MessagesToBeSent)
}