package chatter

func (c *ChatManager) InsertChatterer(chatterer ChatterSendable) {
	c.Lock(); 	c.InsertChatterer_unsafe(chatterer); 	c.Unlock();
}

func (c *ChatManager) SendById(msg ChatMessage) bool {
	c.RLock();	result := c.SendById_unsafe(msg);		c.RUnlock(); return result;
}

func (c *ChatManager) DeleteChatterer(Id string) {
	c.Lock();	c.DeleteChatterer_unsafe(Id);			c.Unlock();
}

func (c *ChatManager) NotifyOnlineUsers_unsafe(notif ChatMessage) {
	for _, chatUser := range(c.ChatUsersByLogin) {
		if(chatUser.IsOnline()) {
			notif.To = chatUser.GetId()
			c.AddChatMessageToChatManagersProcessingQueue(notif)
		}
	}
}

func (c *ChatManager) CreateAndLoginAsChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock()

	if(len(query.Messages) != 2) {
		reply.Message = "ERROR"
	} else {
		_, foundChatUser := c.ChatUsersByLogin[query.Messages[0] + query.Messages[1]]

		if(foundChatUser) {
			reply.Message = "ERROR"
		} else {
			chatUser := NewChatUser(query.Messages[0], query.Messages[1])
			c.InsertChatterer_unsafe(chatUser)

			chatterSendable, foundChatConnection := c.SendToMap[query.OriginConnection]
			chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
		
			if(foundChatConnection && isChatConnection && JoinConnectionToUser(chatConnection, chatUser)) {
				chatConnection.SetNameAndPublicKey(chatUser.GetName(), chatUser.PublicKey)
				reply.Message = GetDetailsAsString(chatUser)
				chatUser.ResendAllPendingMessages()

				if(chatUser.GetChatConnectionCount() == 1) {
					c.NotifyOnlineUsers_unsafe(ChatMessage{OriginConnection: query.OriginConnection, From:"server-new-user-notification",Message:GetDetailsAsString(chatUser)})
				}

			} else {
				reply.Message = "ERROR"
			}

		}
	}

	c.SendById_unsafe(reply)
	c.Unlock()
}

func (c *ChatManager) LoginAsChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock()

	if(len(query.Messages) != 2) {
		reply.Message = "ERROR"
	} else {
		chatterSendable, foundChatConnection := c.SendToMap[query.OriginConnection]
		chatConnection, isChatConnection := chatterSendable.(*ChatConnection)

		chatUser, foundChatUser := c.ChatUsersByLogin[query.Messages[0] + query.Messages[1]]
		if(foundChatConnection && isChatConnection && foundChatUser && JoinConnectionToUser(chatConnection, chatUser)) {
			chatConnection.SetNameAndPublicKey(chatUser.GetName(), chatUser.PublicKey)
			reply.Message = GetDetailsAsString(chatUser)
			chatUser.ResendAllPendingMessages()

			if(chatUser.GetChatConnectionCount() == 1) {
				c.NotifyOnlineUsers_unsafe(ChatMessage{From:"server-new-user-notification",Message:GetDetailsAsString(chatUser)})
			}

		} else {
			reply.Message = "ERROR"
		}
	}

	c.SendById_unsafe(reply)
	c.Unlock()
}

func (c *ChatManager) LogoutFromChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock()

	chatterSendable, foundChatConnection := c.SendToMap[query.From]
	chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
	chatUser := chatConnection.User
	if(foundChatConnection && isChatConnection && chatUser != nil && BreakConnectionFromUser(chatConnection, chatUser)) {
		chatConnection.RemoveNameAndPublicKey()
		reply.Message = chatConnection.GetId()

		if(chatUser.GetChatConnectionCount() == 0) {
			c.NotifyOnlineUsers_unsafe(ChatMessage{From:"server-new-user-notification",Message:GetDetailsAsString(chatUser)})
		}
	} else {
		reply.Message = "ERROR"
	}

	c.SendById_unsafe(reply)
	c.Unlock()
}