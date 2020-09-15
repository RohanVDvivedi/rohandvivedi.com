package chatter

import (
	"sync"
	"time"
)

type Chatterers struct{
	// lock to protect all the chat users
	Lock sync.Mutex

	Chatters map[string]ChatterBox

	InputMessage chan ChatMessage
}

func (c *Chatterers) ChatManagerRun() {
	for msg := range c.InputMessage {

		msgReply := ChatMessage{From: "server", SentAt: time.Now(), To: msg.From, Message: "ERROR"}

		c.Lock.Lock()

		replyToChatterBox, found := c.Chatters[msg.From]

		if(found) {
			switch msg.To {
			case "server-get-chatter-box-name" : {
				chatterBox, found := c.Chatters[msg.Message]
				if(found) {
					msgReply.Message = chatterBox.GetName()
				}
			}
			case "server-get-chat-user-publickey" : {
				chatterBox, found := c.Chatters[msg.Message]
				if(found) {
					chatUser, isChatUser := chatterBox.(*ChatUser)
					if(isChatUser) {
						msgReply.Message = chatUser.PublicKey
					}
				}
			}
			case "server-create-chat-group" : {
			}
			case "server-add-user-to-chat-group" : {
			}
			case "server-delete-chat-group" : {
			}
			}
		}
		c.Lock.Unlock()

		if(found) {
			go replyToChatterBox.SendMessage(msgReply)
		}

	}
}

func (c *Chatterers) InsertChatterBox(chatterBox ChatterBox) {
	c.Lock.Lock()
	c.Chatters[chatterBox.GetId()] = chatterBox
	c.Lock.Unlock()
}

func (c *Chatterers) DeleteChatterBox(Id string) {
	c.Lock.Lock()
	chatterBox, found := c.Chatters[Id]
	if(found) {
		chatterBox.Destroy()
		delete(c.Chatters, Id);
	}
	c.Lock.Unlock()
}

func (c *Chatterers) GetChatterBoxById(Id string) ChatterBox {
	c.Lock.Lock()
	chatterBox, found := c.Chatters[Id]
	c.Lock.Unlock()
	if(found) {
		return chatterBox;
	}
	return nil
}
