package chatter

import (
	"sync"
	"time"
)

type ChatManager struct{
	// lock to protect all the chat users
	Lock sync.Mutex

	// Chatters => id => ChatterSendables (chat groups and users and even connections)
	Chatters map[string]ChatterSendable

	// Chattereres => name => id => ChatterBox (chat groups and users)
	Chatterers map[string]map[string]ChatterBox

	// this is for authentication purpose
	// chat users mapped with name+publickey -> chat user
	ChatUsersMapped map[string]*ChatUser

	ServerMessagesToBeProcessed *ChatMessageQueue
}

func NewChatManager() *ChatManager {
	cm := &ChatManager{
		Chatters: make(map[string]ChatterSendable),
		Chatterers: make(map[string]map[string]ChatterBox),
		ServerMessagesToBeProcessed: NewChatMessageQueue(),
	}
	return cm
}

func (c *ChatManager) ChatManagerRun() {
	for (true) {
		msg := c.ServerMessagesToBeProcessed.Top()

		msgReply := ChatMessage{From: "server", SentAt: time.Now(), To: msg.From, Message: "ERROR"}

		c.Lock.Lock()

		replyToChatterBox, found := c.Chatters[msg.From]

		if(found) {
			switch msg.To {
			case "server-get-chatter-box-name" : {
				chatterSendable, found := c.Chatters[msg.Message]
				if(found) {
					chatterBox, isChatterBox := chatterSendable.(ChatterBox)
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
				chatterSendable, foundChatConnection := c.Chatters[msg.From]
				chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
				chatUser, foundChatUser := c.ChatUsersMapped[msg.Message]
				if(foundChatConnection && isChatConnection && foundChatUser) {
					chatUser.AddChatConnection(chatConnection)
					chatConnection.SetChatUser(chatUser)
					msgReply.To = chatUser.GetId()
					msgReply.Message = "Logged in with " + chatConnection.GetId()
				}
			}
			case "server-logout" : {

			}
			case "server-add-user-to-chat-group" : {
			}
			case "server-delete-chat-group" : {
			}
			case "server-delete-chat-user" : {
			}
			case "server-notify-all" : {
			}
			}
		}
		c.Lock.Unlock()

		if(found) {
			go replyToChatterBox.SendMessage(msgReply)
		}

		c.ServerMessagesToBeProcessed.Top()
	}
}

func (c *ChatManager) InsertChatterer(chatterer ChatterSendable) {
	c.Lock.Lock()

	// insert to the main map allowing us to send messages
	c.Chatters[chatterer.GetId()] = chatterer

	// insertions by name for groups and users for find by name stuff
	chatterBox, isChatterBox := chatterer.(ChatterBox)
	if(isChatterBox) {
		_, chatterBoxesPresent := c.Chatterers[chatterBox.GetName()]
		if(!chatterBoxesPresent) {
			c.Chatterers[chatterBox.GetName()] = make(map[string]ChatterBox)
		}
		c.Chatterers[chatterBox.GetName()][chatterBox.GetId()] = chatterBox
	}

	// insertion by name and public key for authentication
	chatUser, isChatUser:= chatterer.(*ChatUser)
	if(isChatUser) {
		c.ChatUsersMapped[chatUser.GetName() + "," + chatUser.PublicKey] = chatUser
	}

	chatterer.SendMessage(ChatMessage{From:"server",To:chatterer.GetId(),SentAt:time.Now(),Message:"Chatterer registered"})
	c.Lock.Unlock()
}

func (c *ChatManager) DeleteChatterer(Id string) {
	c.Lock.Lock()
	chatterSendable, found := c.Chatters[Id]
	if(found) {

		delete(c.Chatters, Id);

		chatterBox, isChatterBox := chatterSendable.(ChatterBox)
		if(isChatterBox) {
			delete(c.Chatterers[chatterBox.GetName()], chatterBox.GetId());
		}

		chatUser, isChatUser := chatterSendable.(*ChatUser)
		if(isChatUser) {
			delete(c.ChatUsersMapped, chatUser.GetName() + "," + chatUser.PublicKey);
		}

		chatterSendable.Destroy()
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
