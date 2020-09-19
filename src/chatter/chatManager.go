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
		// All the chat connections, users and groups mapped together by their connection ids
		SendToMap: make(map[string]ChatterSendable),

		// chat users and groups mapped by name and then by id 
		// UsersAndGroups[name][id] => ChatterBox, used for search by name
		UsersAndGroups: make(map[string]map[string]ChatterBox),

		// chat users mapped by login credentials, i.e. string(name + publicKey)
		ChatUsersByLogin: make(map[string]*ChatUser),

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

func (c *ChatManager) ChatManagerProcessServerRequests() {
	for (true) {
		msg := c.ServerMessagesToBeProcessed.Top()
		c.ServerMessagesToBeProcessed.Pop()

		if(msg.IsValidServerRequest()) {
			switch msg.To {
				// no params required
				// huge queries allowed allowed to all users
				case "server-get-all-users" : {
					c.GetAllUsers(msg)
				}
				case "server-get-all-groups" : {
					c.GetAllGroups(msg)
				}
				case "server-get-all-online-users" : {
					c.GetAllOnlineUsers(msg)
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
			}
		} else if(msg.IsValidChatMessage()) {
			c.SendById(msg)
		}
	}
}
