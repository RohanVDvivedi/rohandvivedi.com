package chatter

import (
	"sync"
	"golang.org/x/net/websocket"
)

type Chatterers struct{
	// lock to protect all the chat users
	Lock sync.Mutex

	Chatters map[string]ChatterBox
}

// returns pointer to the created chat user is the user was created
func (c *Chatterers) InsertUniqueChatUserByName(name string, conn *websocket.Conn) *ChatUser {
	var chatUser *ChatUser = nil
	c.Lock.Lock()
	_, chatUserSameNameExists := c.Chatters[name]
	if(!chatUserSameNameExists) {
		chatUser = NewChatUser(name, conn)
		c.Chatters[chatUser.Name] = chatUser
	}
	c.Lock.Unlock()
	return chatUser
}

// returns true if the user is removed
func (c *Chatterers) DeleteChatterBoxByName(name string) bool {
	c.Lock.Lock()
	chatterBox, found := c.Chatters[name]
	if(found) {
		chatUser, isChatUser := chatterBox.(*ChatUser)
		if(isChatUser) {
			chatUser.DestroyChatUser()
			delete(c.Chatters, name);
		}
	}
	c.Lock.Unlock()
	return found
}

func (c *Chatterers) GetChatUserByName(name string) ChatterBox {
	c.Lock.Lock()
	chatterBox, found := c.Chatters[name]
	c.Lock.Unlock()
	if(found) {
		return chatterBox;
	}
	return nil
}
