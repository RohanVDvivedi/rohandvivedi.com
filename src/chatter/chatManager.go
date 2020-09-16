package chatter

import (
	"sync"
	"time"
)

type ChatManager struct{
	// lock to protect all the chat users
	Lock sync.Mutex

	Chatters map[string]ChatterSendable

	ServerMessagesToBeProcessed* ChatMessageQueue
}

func (c *ChatManager) ChatManagerRun() {
	for msg := range c.InputMessage {

		msgReply := ChatMessage{From: "server", SentAt: time.Now(), To: msg.From, Message: "ERROR"}

		c.Lock.Lock()

		replyToChatterBox, found := c.Chatters[msg.From]

		if(found) {
			switch msg.To {
			case "server-get-chatter-box-name" : {
				chatterSendable, found := c.Chatters[msg.Message]
				if(found) {
					chatterBox isChatterBox := chatterSendable.(ChatterBox)
					if(isChatterBox) {
						msgReply.Message = chatterBox.GetName()
					}
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
			case "server-create-chat-user" : {
			}
			case "server-login-as-chat-user" : {
			}
			case "server-add-user-to-chat-group" : {
			}
			case "server-delete-chat-group" : {
			}
			case "server-notify-all" : {
			}
			}
		}
		c.Lock.Unlock()

		if(found) {
			go replyToChatterBox.SendMessage(msgReply)
		}

	}
}

func (c *ChatManager) InsertChatterer(chatterer ChatterSendable) {
	c.Lock.Lock()
	c.Chatters[chatterBox.GetId()] = chatterBox
	c.Lock.Unlock()
}

func (c *ChatManager) DeleteChatterer(Id string) {
	c.Lock.Lock()
	chatterBox, found := c.Chatters[Id]
	if(found) {
		chatterBox.Destroy()
		delete(c.Chatters, Id);
	}
	c.Lock.Unlock()
}

func (c *ChatManager) SendById(msg ChatMessage) bool {
	c.Lock.Lock()
	chatterSendable, found := c.Chatters[msg.To]
	if(found) {
		chatterSendable.SendMessage(msg);
	}
	c.Lock.Unlock()
	return found
}
