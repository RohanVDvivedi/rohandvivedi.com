package chatter

func (c *ChatManager) InsertChatterer(chatterer ChatterSendable) {
	c.Lock.Lock(); 	c.InsertChatterer_unsafe(chatterer); 	c.Lock.Unlock();
}

func (c *ChatManager) SendById(msg ChatMessage) bool {
	c.Lock.Lock();	result := c.SendById_unsafe(msg);		c.Lock.Unlock(); return result;
}

func (c *ChatManager) DeleteChatterer(Id string) {
	c.Lock.Lock();	c.DeleteChatterer_unsafe(Id);			c.Lock.Unlock();
}

func (c *ChatManager) CreateAndLoginAsChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

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
				reply.Message = chatUser.GetId()
				chatUser.ResendAllPendingMessages()
			} else {
				reply.Message = "ERROR"
			}

		}
	}

	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) LoginAsChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	if(len(query.Messages) != 2) {
		reply.Message = "ERROR"
	} else {
		chatterSendable, foundChatConnection := c.SendToMap[query.OriginConnection]
		chatConnection, isChatConnection := chatterSendable.(*ChatConnection)

		chatUser, foundChatUser := c.ChatUsersByLogin[query.Messages[0] + query.Messages[1]]
		if(foundChatConnection && isChatConnection && foundChatUser && JoinConnectionToUser(chatConnection, chatUser)) {
			chatConnection.SetNameAndPublicKey(chatUser.GetName(), chatUser.PublicKey)
			reply.Message = chatUser.GetId()
			chatUser.ResendAllPendingMessages()
		} else {
			reply.Message = "ERROR"
		}
	}

	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) LogoutAllConnectionsFromChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	chatterSendable, foundChatConnection := c.SendToMap[query.From]
	chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
	if(foundChatConnection && isChatConnection && chatConnection.User != nil && BreakConnectionFromUser(chatConnection, chatConnection.User)) {
		chatConnection.RemoveNameAndPublicKey()
		reply.Message = chatConnection.GetId()
	} else {
		reply.Message = "ERROR"
	}

	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}