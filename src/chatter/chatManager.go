package chatter

import (
	"sync"
)

type ChatManager struct{
	sync.RWMutex

	// All the chat connections, users and groups mapped together by their connection ids
	SendToMap map[string]ChatterSendable

	// chat users and groups mapped by name and then by id 
	// UsersAndGroups[name][id] => ChatterBox, used for search by name
	UsersAndGroups map[string]map[string]ChatterBox

	// chat users mapped by login credentials, i.e. string(name + publicKey)
	ChatUsersByLogin map[string]*ChatUser

	ServerMessagesToBeProcessed *ChatMessageQueue
}

func NewChatManager() *ChatManager {
	cm := &ChatManager{
		SendToMap: make(map[string]ChatterSendable),
		UsersAndGroups: make(map[string]map[string]ChatterBox),
		ChatUsersByLogin: make(map[string]*ChatUser),
		ServerMessagesToBeProcessed: NewChatMessageQueue(),
	}
	go cm.ChatManagerProcessServerRequests()
	return cm
}

func (c *ChatManager) AddChatMessageToChatManagersProcessingQueue(msg ChatMessage) {
	if(msg.IsValidChatMessage() || msg.IsValidServerRequest() || msg.IsValidServerResponse()) {
		c.ServerMessagesToBeProcessed.Push(msg)
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
					c.GetAllMyGroups(msg)
				}
				case "server-get-all-my-active-connections" : {
					c.GetAllMyActiveConnections(msg)
				}

				// 1 params required in the message, which must be equal to id or name
				// Message: Id or name of some one whose details you want
				case "server-search-chatter-box" : {
					c.SearchChatterBox(msg)
				}
				
				// create, login, logout and delete user calls, these must be infrequent
				// Message : contains name,publicKey to create a corresponding user
				// Message must be comming from a chat connection only
				case "server-create-and-login-as-chat-user" : {
					c.CreateAndLoginAsChatUser(msg)
				}
				case "server-login-as-chat-user" : {
					c.LoginAsChatUser(msg)
				}
				case "server-logout-from-chat-user" : {
					c.LogoutFromChatUser(msg)
				}
			}
		} else if(msg.IsValidChatMessage() || msg.IsValidServerResponse()) {
			c.SendById(msg)
		}
	}
}
