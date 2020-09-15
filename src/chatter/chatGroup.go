package chatter

import (
	"time"
)

type ChatGroup struct {
	Id string
	Name string

	MessagesToBeSent chan ChatMessage

	ChatUsers map[string]*ChatUser

	LastMessage time.Time
}

func NewChatGroup(name string) *ChatGroup {
	grp := &ChatGroup{
		Id: GetNewChatGroupId(),
		Name: name,
		MessagesToBeSent: make(chan ChatMessage, 10),
		ChatUsers: make(map[string]*ChatUser),
		LastMessage: time.Now(),
	}
	go grp.LoopOverChannelToPassMessages()
	return grp
}

func (grp *ChatGroup) GetId() string {
	return grp.Id
}

func (grp *ChatGroup) GetName() string {
	return grp.Name
}

func (grp *ChatGroup) SendMessage(msg ChatMessage) {
	grp.MessagesToBeSent <- msg
}

func (grp *ChatGroup) LoopOverChannelToPassMessages() {
	for msg := range grp.MessagesToBeSent {
		if(msg.To == grp.Id) {
			for _, user := range grp.ChatUsers {
				go user.SendMessage(msg)
			}
			grp.LastMessage = time.Now()
		}
	}
}

func (grp *ChatGroup) Destroy() {
	close(grp.MessagesToBeSent)
}