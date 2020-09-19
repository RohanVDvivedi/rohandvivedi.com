package chatter

func (c *ChatManager) InsertChatterer(chatterer ChatterSendable) {
	c.Lock.Lock() 	c.InsertChatterer_unsafe(chatterer) 	c.Lock.Unlock()
}

func (c *ChatManager) SendById_unsafe(msg ChatMessage) bool {
	c.Lock.Lock()	result := c.SendById_unsafe(chatterer)	c.Lock.Unlock(); return result;
}

func (c *ChatManager) DeleteChatterer(Id string) {
	c.Lock.Lock()	c.DeleteChatterer_unsafe(chatterer)		c.Lock.Unlock()
}

func (c *ChatManager) CreateAndLoginAsChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

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

	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) LoginAsChatUser(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	if(len(msg.Messages) != 2) {
		reply.Message = "ERROR"
	} else {
		chatterSendable, foundChatConnection := c.Chatters[msg.OriginConnection]
		chatConnection, isChatConnection := chatterSendable.(*ChatConnection)

		chatUser, foundChatUser := c.ChatUsersMapped[msg.Message]
		if(foundChatConnection && isChatConnection && foundChatUser && JoinConnectionToUser(chatConnection, chatUser)) {
			chatConnection.SetNameAndPublicKey(chatUser.GetName(), chatUser.PublicKey)
			reply.Message = chatUser.GetId()
			chatUser.ResendAllPendingMessages()
		} else if (foundChatConnection && isChatConnection) {
			reply.Message = "ERROR"
		}

	}

	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) LogoutAllConnectionsFromChatUser(query ChatMessage) {
	c.Lock.Lock()

	chatterSendable, foundChatConnection := c.Chatters[msg.From]
	chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
	if(foundChatConnection && isChatConnection && chatConnection.User != nil && BreakConnectionFromUser(chatConnection, chatConnection.User)) {
		chatConnection.RemoveNameAndPublicKey()
		reply.Message = chatConnection.GetId()
		c.SendById_unsafe(reply)
	} else {
		reply := StdReplyToOrigin(query)
		reply.Message = "ERROR"
		c.SendById_unsafe(reply)
	}

	c.Lock.Unlock()
}