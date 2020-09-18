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
		Chatters: make(map[string]ChatterSendable),
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

func (c *ChatManager) ChatManagerProcessServerRequests() {
	for (true) {
		msg := c.ServerMessagesToBeProcessed.Top()
		c.ServerMessagesToBeProcessed.Pop()

		c.Lock.Lock()

		if(msg.IsValidServerRequest()) {
			serverReplies := []ChatMessage{}
			//stdReplyFrom := ChatMessage{From:msg.To, To:msg.From 			  , ContextId: msg.MessageId}
			stdReplyOrigin := ChatMessage{From:msg.To, To:msg.OriginConnection, ContextId: msg.MessageId}

			switch msg.To {
				// returns id, name and public key of all users
				case "server-get-all-users" : {
					reply := stdReplyOrigin
					reply.Messages = []string{}
					if(IsChatUserId(msg.From)) {
						for _, chatUser := range(c.ChatUsersMapped) {
							reply.Messages = append(reply.Messages, chatUser.GetId() + "," + chatUser.GetName() + "," + strconv.Itoa(chatUser.GetChatConnectionCount()))
						}
					} else {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
				}
				// returns id, name and public key of all users
				case "server-get-all-online-users" : {
					reply := stdReplyOrigin
					reply.Messages = []string{}
					if(IsChatUserId(msg.From)) {
						for _, chatUser := range(c.ChatUsersMapped) {
							if(chatUser.GetChatConnectionCount() > 0) {
								reply.Messages = append(reply.Messages, chatUser.GetId() + "," + chatUser.GetName()
							}
						}
					} else {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
				}
				// returns id, name of all groups of the corresponding user
				case "server-get-all-my-groups" : {
					reply := stdReplyOrigin
					reply.Messages = []string{}
					chatterSendable, found := c.Chatters[msg.From]
					chatUser, isChatUser := chatterSendable.(*ChatUser)
					if(found && isChatUser) {
						for _, chatGroup := range(chatUser.ChatGroups) {
							reply.Messages = append(reply.Messages, chatGroup.GetId() + "," + chatGroup.GetName())
						}
					} else {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
				}
				// returns id, name of all groups of the corresponding user
				case "server-get-all-my-active-connections" : {
					reply := stdReplyOrigin
					reply.Messages = []string{}
					chatterSendable, found := c.Chatters[msg.From]
					chatUser, isChatUser := chatterSendable.(*ChatUser)
					if(found && isChatUser) {
						for _, chatConnection := range(chatUser.ChatConnections) {
							reply.Messages = append(reply.Messages, chatConnection.GetId())
						}
					} else {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
				}
				// Message: Id of some one whose name is to be found
				case "server-get-chatter-box-name" : {
					reply := stdReplyOrigin
					chatterSendable, found := c.Chatters[msg.Message]
					chatterBox, isChatterBox := chatterSendable.(ChatterBox)
					if(found && isChatterBox) {
						reply.Message = chatterBox.GetName()
					} else {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
				}
				// Message: Id of some one whose public key is to be found
				case "server-get-chat-user-publickey" : {
					reply := stdReplyOrigin
					chatterBox, found := c.Chatters[msg.Message]
					chatUser, isChatUser := chatterBox.(*ChatUser)
					if(found && isChatUser) {
						reply.Message = chatUser.PublicKey
					} else {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
				}
				// Message : contains name,publicKey to create a corresponding user
				case "server-create-and-login-as-chat-user" : {
					reply := stdReplyOrigin
					chatterSendable, foundChatConnection := c.Chatters[msg.OriginConnection]
					chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
					_, foundChatUser := c.ChatUsersMapped[msg.Message]
					params := strings.Split(msg.Message, ",")
					if(foundChatConnection && isChatConnection && !foundChatUser && len(params) == 2) {
						chatUser := NewChatUser(params[0], params[1])
						c.InsertChatterer_unsafe(chatUser)
						JoinConnectionToUser(chatConnection, chatUser)
						chatConnection.SetNameAndPublicKey(chatUser.GetName(), chatUser.PublicKey)
						reply.Message = chatUser.GetId()
					} else if (foundChatConnection && isChatConnection) {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
				}
				// Message : contains name,publicKey to login to that user
				case "server-login-as-chat-user" : {
					reply := stdReplyOrigin
					chatterSendable, foundChatConnection := c.Chatters[msg.OriginConnection]
					chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
					chatUser, foundChatUser := c.ChatUsersMapped[msg.Message]
					if(foundChatConnection && isChatConnection && foundChatUser && JoinConnectionToUser(chatConnection, chatUser)) {
						chatConnection.SetNameAndPublicKey(chatUser.GetName(), chatUser.PublicKey)
						reply.Message = chatUser.GetId()
					} else if (foundChatConnection && isChatConnection) {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
				}
				case "server-logout" : {
					reply := stdReplyOrigin
					chatterSendable, foundChatConnection := c.Chatters[msg.From]
					chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
					if(foundChatConnection && isChatConnection && chatConnection.User != nil && BreakConnectionFromUser(chatConnection, chatConnection.User)) {
						chatConnection.RemoveNameAndPublicKey()
						reply.Message = chatConnection.GetId()
					} else {
						reply.Message = "ERROR"
					}
					serverReplies = append(serverReplies, reply)
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
