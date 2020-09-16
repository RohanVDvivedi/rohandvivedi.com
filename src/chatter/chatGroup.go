package chatter

import (
	"time"
)

type ChatGroup struct {
	Id
	Name

	MessagesToBeSent ChatMessageQueue

	ChatUsers map[string]*ChatUser
}

func NewChatGroup(name string) *ChatGroup {
	grp := &ChatGroup{
		Id: Id{GetNewChatUserId()},
		Name: Name{name},
		MessagesToBeSent: NewChatMessageQueue(),
		ChatUsers: make(map[string]*ChatUser),
	}
	go grp.LoopToForwardMessages()
	return grp
}

func (grp *ChatGroup) SendMessage(msg ChatMessage) {
	grp.MessagesToBeSent.Push(msg)
}

func (grp *ChatGroup) LoopToForwardMessages() {
	for (true) {
		msg := grp.MessagesToBeSent.Top()
		if(msg.To == grp.GetId()) {
			for _, user := range grp.ChatUsers {
				user.SendMessage(msg)
			}
		}
		grp.MessagesToBeSent.Pop()
	}
}

func (grp *ChatGroup) Destroy() {
	close(grp.MessagesToBeSent)
}