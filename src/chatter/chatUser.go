package chatter

import (
	"fmt"
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
		MessagesPendingToBeSent: NewChatMessageQueue(),
		ChatConnections: make(map[string]*ChatConnection),
		ChatGroups: make(map[string]*ChatGroup),
	}
	return user
}

func (user *ChatUser) SendMessage(msg ChatMessage) error {
	_, usersGroup := user.ChatGroups[msg.To]
	if(msg.To == user.GetId() || usersGroup) {
		sentTo := 0
		for _, cconn := range user.ChatConnections {
			err := cconn.SendMessage(msg)
			if(err == nil) {
				sentTo += 1
			} else {
				fmt.Println(err)
				cconn.Destroy()
			}
		}
		if(sentTo == 0) {
			user.MessagesPendingToBeSent.Push(msg)
		}
	}
	return nil
}

func (user *ChatUser) Destroy() {
	user.MessagesPendingToBeSent = nil
}

/* Joinery methods */
func (user *ChatUser) HasChatConnection(c *ChatConnection) bool {
	_, found := user.ChatConnections[c.GetId()]
	return found
}
func (user *ChatUser) AddChatConnection(c *ChatConnection) {
	user.ChatConnections[c.GetId()] = c
}
func (user *ChatUser) RemoveChatConnection(c *ChatConnection) {
	delete(user.ChatConnections, c.GetId())
}
func (user *ChatUser) GetChatConnectionCount() int {
	return len(user.ChatConnections)
}

func (user *ChatUser) HasChatGroup(c *ChatGroup) bool {
	_, found := user.ChatGroups[c.GetId()]
	return found
}
func (user *ChatUser) AddChatGroup(c *ChatGroup) {
	user.ChatGroups[c.GetId()] = c
}
func (user *ChatUser) RemoveChatGroup(c *ChatGroup) {
	delete(user.ChatGroups, c.GetId())
}
func (user *ChatUser) GetChatGroupCount() int {
	return len(user.ChatGroups)
}

/* Below methods update modify session values, of all the chat connections of this chat user */
func (user *ChatUser) GetNameAndPublicKey() (name string, publicKey string) {
	return user.GetName(), user.PublicKey
}

func (user *ChatUser) SetNameAndPublicKey(name string, publicKey string) {
	user.SetName(name)
	user.PublicKey = publicKey
	for _, cconn := range user.ChatConnections {
		cconn.SetNameAndPublicKey(name, publicKey)
	}
}