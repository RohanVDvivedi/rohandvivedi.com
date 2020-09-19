package chatter

import (
	"sync"
	"time"
	"strings"
	"strconv"
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
		SendToMap: make(map[string]ChatterSendable),
		Chatterers: make(map[string]map[string]ChatterBox),
		ChatUsersMapped: make(map[string]*ChatUser),
		ServerMessagesToBeProcessed: NewChatMessageQueue(),
	}
	go cm.ChatManagerProcessServerRequests()
	return cm
}

func (c *ChatManager) AddChatMessageToChatManagersProcessingQueue(msg ChatMessage) {
	if(msg.IsValidChatMessage() || msg.IsValidServerRequest()) {
		Chatters.ServerMessagesToBeProcessed.Push(msg)
	}
}

func StdReplyToOrigin(ChatMessage msg) ChatMessage {
	return ChatMessage{From:msg.To, To:msg.OriginConnection, ContextId: msg.MessageId, Message: "", Messages: []string{}}
}

func StdReplyToSender(ChatMessage msg) ChatMessage {
	return ChatMessage{From:msg.To, To:msg.From, ContextId: msg.MessageId, Message: "", Messages: []string{}}
}

func GetDetailsString(cs ChatterSendable) string {
	return ""
}

func (c *ChatManager) ChatManagerProcessServerRequests() {
	for (true) {
		msg := c.ServerMessagesToBeProcessed.Top()
		c.ServerMessagesToBeProcessed.Pop()

		c.Lock.Lock()

		if(msg.IsValidServerRequest()) {
			serverReplies := []ChatMessage{}
			
			switch msg.To {
				// no params required
				// huge queries allowed allowed to all users
				case "server-get-all-users" : {
					serverReplies = append(serverReplies, c.GetAllUsers(msg))
				}
				case "server-get-all-groups" : {
					serverReplies = append(serverReplies, c.GetAllGroups(msg))
				}
				case "server-get-all-online-users" : {
					serverReplies = append(serverReplies, c.GetAllOnlineUsers(msg))
				}

				// no params required
				// mid sized queries allowed to users to know their own groups and active connections
				case "server-get-all-my-groups" : {
					serverReplies = append(serverReplies, c.GetAllMyGroups(msg))
				}
				case "server-get-all-my-active-connections" : {
					serverReplies = append(serverReplies, c.GetAllMyActiveConnections(msg))
				}

				// 1 params required in the message, which must be equal to id or name
				// Message: Id or name of some one whose details you want
				case "server-search-chatter-box" : {
					serverReplies = append(serverReplies, c.SearchChatterBox(msg))
				}
				
				// create, login, logout and delete user calls, these must be infrequent
				// Message : contains name,publicKey to create a corresponding user
				// Message must be comming from a chat connection only
				case "server-create-and-login-as-chat-user" : {
					serverReplies = append(serverReplies, c.CreateAndLoginAsChatUser(msg))
				}
				case "server-login-as-chat-user" : {
					serverReplies = append(serverReplies, c.LoginAsChatUser(msg))
				}
				case "server-logout-current-session-from-chat-user" : {
					serverReplies = append(serverReplies, c.LogoutCurrentCessionFromChatUser(msg))
				}

				case "server-create-chat-group" : {
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

			for _, msg := range serverReplies {
				c.SendById_unsafe(msg)
			}
		} else if(msg.IsValidChatMessage()) {
			c.SendById_unsafe(msg)
		}

		c.Lock.Unlock()
	}
}

func (c *ChatManager) InsertChatterer_unsafe(chatterer ChatterSendable) {
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

	chatterer.SendMessage(ChatMessage{From:"server-chatterer-created",To:chatterer.GetId(),SentAt:time.Now(),Message:"Chatterer registered"})
}

func (c *ChatManager) InsertChatterer(chatterer ChatterSendable) {
	c.Lock.Lock()
	c.InsertChatterer_unsafe(chatterer)
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

func (c *ChatManager) SendById_unsafe(msg ChatMessage) bool {
	chatterSendable, found := c.Chatters[msg.To]
	if(found) {
		chatterSendable.SendMessage(msg);
	}
	return found
}
