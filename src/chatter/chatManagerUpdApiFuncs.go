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
	errorOccurred := true
	c.Lock()

	if(len(query.Messages) == 2) {
		_, foundChatUser := c.ChatUsersByLogin[query.Messages[0] + query.Messages[1]]

		if(!foundChatUser) {
			chatUser := NewChatUser(query.Messages[0], query.Messages[1])
			c.InsertChatterer_unsafe(chatUser)

			chatterSendable, foundChatConnection := c.SendToMap[query.OriginConnection]
			chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
		
			if(foundChatConnection && isChatConnection && JoinConnectionToUser(chatConnection, chatUser)) {
				chatConnection.SetNameAndPublicKey(chatUser.GetName(), chatUser.PublicKey)
				reply.Message = chatUser.GetDetailsAsString()

				c.SendById_unsafe(reply)
				chatUser.ResendAllPendingMessages()
				c.NotifyOnlineUsers_unsafe(ChatMessage{OriginConnection: query.OriginConnection,From:"server-event-update",Message:chatUser.GetDetailsAsString()})

				errorOccurred = false
			}
		}
	}

	if(errorOccurred) {
		reply.Message = "ERROR"
		c.SendById_unsafe(reply)
	}

	c.SendById_unsafe(reply)
	c.Unlock()
}

func (c *ChatManager) LoginAsChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	errorOccurred := true
	c.Lock()

	if(len(query.Messages) == 2) {
		chatterSendable, foundChatConnection := c.SendToMap[query.OriginConnection]
		chatConnection, isChatConnection := chatterSendable.(*ChatConnection)

		chatUser, foundChatUser := c.ChatUsersByLogin[query.Messages[0] + query.Messages[1]]
		if(foundChatConnection && isChatConnection && foundChatUser && JoinConnectionToUser(chatConnection, chatUser)) {
			chatConnection.SetNameAndPublicKey(chatUser.GetName(), chatUser.PublicKey)
			reply.Message = chatUser.GetDetailsAsString()

			c.SendById_unsafe(reply)
			chatUser.ResendAllPendingMessages()
			c.NotifyOnlineUsers_unsafe(ChatMessage{OriginConnection: query.OriginConnection, From:"server-event-update",Message:chatUser.GetDetailsAsString()})

			errorOccurred = false
		}
	}

	if(errorOccurred) {
		reply.Message = "ERROR"
		c.SendById_unsafe(reply)
	}

	c.SendById_unsafe(reply)
	c.Unlock()
}

func (c *ChatManager) LogoutFromChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	errorOccurred := true
	c.Lock()

	chatterSendable, foundChatConnection := c.SendToMap[query.From]
	chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
	chatUser := chatConnection.User
	if(foundChatConnection && isChatConnection && chatUser != nil && BreakConnectionFromUser(chatConnection, chatUser)) {
		chatConnection.RemoveNameAndPublicKey()
		reply.Message = chatConnection.GetDetailsAsString()

		c.SendById_unsafe(reply)
		c.NotifyOnlineUsers_unsafe(ChatMessage{OriginConnection: query.OriginConnection, From:"server-event-update",Message:chatUser.GetDetailsAsString()})

		errorOccurred = false
	}

	if(errorOccurred) {
		reply.Message = "ERROR"
		c.SendById_unsafe(reply)
	}

	c.Unlock()
}

func (c *ChatManager) LogoutFromAndDeleteChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	errorOccurred := true
	c.Lock()

	chatterSendable, foundChatConnection := c.SendToMap[query.From]
	chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
	chatUser := chatConnection.User
	if(foundChatConnection && isChatConnection && chatUser != nil) {
		
		for _, CUConn := range(chatUser.ChatConnections) {
			if(BreakConnectionFromUser(chatConnection, CUConn)) {
				CUConn.RemoveNameAndPublicKey()
			}
		}

		reply.Message = chatConnection.GetDetailsAsString()
		c.SendById_unsafe(reply)

		reply = StdReplyToFrom(query)
		c.SendById_unsafe(reply)

		c.NotifyOnlineUsers_unsafe(ChatMessage{OriginConnection: query.OriginConnection, From:"server-event-update",Message:chatUser.GetDetailsAsString()})

		errorOccurred = false
	}

	if(errorOccurred) {
		reply.Message = "ERROR"
		c.SendById_unsafe(reply)
	}

	c.Unlock()
}