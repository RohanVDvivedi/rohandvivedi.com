package chatter

import (
	"sync"
	"fmt"
)

type ChatUser struct {
	sync.RWMutex

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

// a chat user is online if user has atleast 1 active chat connection
func (user *ChatUser) IsOnline() bool {
	return len(user.ChatConnections) > 0
}

func (user *ChatUser) ResendAllPendingMessages() {
	oldMsgs := user.MessagesPendingToBeSent
	count := oldMsgs.MessageCount()
	for(count > 0) {
		user.SendMessage(oldMsgs.Top())
		oldMsgs.Pop()
		count--
	}
	user.MessagesPendingToBeSent = NewChatMessageQueue()
}

// message must be sent to the user or one of the groups that the user is member of
// message must be sent by the server, or any chat user or 
func (user *ChatUser) SendMessage(msg ChatMessage) error {
	_, msgIsFromUsersConnections := user.ChatConnections[msg.From]
	_, msgIsToAUsersGroupThatThisUserIsAPartOf := user.ChatGroups[msg.To]
	if((msgIsFromUsersConnections || IsChatUserId(msg.From) || IsChatManagerId(msg.From)) && 
	(msg.To == user.GetId() || msgIsToAUsersGroupThatThisUserIsAPartOf)) {
		sentTo := 0
		for _, cconn := range user.ChatConnections {
			err := cconn.SendMessage(msg)
			if(err == nil) {
				sentTo += 1
			} else {
				fmt.Println(err)
				BreakConnectionFromUser(cconn, user)
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