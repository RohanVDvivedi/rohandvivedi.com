package chatter

func (c *ChatManager) InsertChatterer(chatterer ChatterSendable) {
	c.Lock(); 	c.InsertChatterer_unsafe(chatterer); 	c.Unlock();
}

func (c *ChatManager) SendById(msg ChatMessage) bool {
	c.Lock();	result := c.SendById_unsafe(msg);		c.Unlock(); return result;
}

func (c *ChatManager) DeleteChatterer(Id string) {
	c.Lock();	c.DeleteChatterer_unsafe(Id);			c.Unlock();
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
				reply.Message = chatUser.GetDetailsAsString()
				chatUser.ResendAllPendingMessages()

				if(chatUser.GetChatConnectionCount() == 1) {
					c.NotifyOnlineUsers_unsafe(ChatMessage{OriginConnection: query.OriginConnection, From:"server-new-user-notification",Message:chatUser.GetDetailsAsString()})
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
			reply.Message = chatUser.GetDetailsAsString()
			chatUser.ResendAllPendingMessages()

			if(chatUser.GetChatConnectionCount() == 1) {
				c.NotifyOnlineUsers_unsafe(ChatMessage{From:"server-new-user-notification",Message:chatUser.GetDetailsAsString()})
			}

		} else {
			reply.Message = "ERROR"
		}
	}

	c.SendById_unsafe(reply)
	c.Unlock()
}

func (c *ChatManager) LogoutAllConnectionsFromChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock()

	chatterSendable, foundChatConnection := c.SendToMap[query.From]
	chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
	chatUser := chatConnection.User
	if(foundChatConnection && isChatConnection && chatUser != nil && BreakConnectionFromUser(chatConnection, chatUser)) {
		chatConnection.RemoveNameAndPublicKey()
		reply.Message = chatConnection.GetDetailsAsString()

		if(chatUser.GetChatConnectionCount() == 0) {
			c.NotifyOnlineUsers_unsafe(ChatMessage{From:"server-new-user-notification",Message:chatUser.GetDetailsAsString()})
		}
	} else {
		reply.Message = "ERROR"
	}

	c.SendById_unsafe(reply)
	c.Unlock()
}