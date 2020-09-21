package chatter

// all the below queries must be called only after login and sent with the user_id as from field

func (c *ChatManager) GetAllUsers(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	if(IsChatUserId(query.From)) {
		for _, chatUser := range(c.ChatUsersByLogin) {
			reply.Messages = append(reply.Messages, chatUser.GetDetailsAsString())
		}
	} else {
		reply.Message = "ERROR"
	}

	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) GetAllGroups(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	if(IsChatUserId(query.From)) {
		for _, chatterBoxesByName := range(c.UsersAndGroups) {
			for _, chatterBox := range(chatterBoxesByName) {
				_, isChatGroup := chatterBox.(*ChatGroup)
				if(isChatGroup) {
					reply.Messages = append(reply.Messages, chatterBox.GetDetailsAsString())
				}
			}
		}
	} else {
		reply.Message = "ERROR"
	}

	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) GetAllOnlineUsers(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	if(IsChatUserId(query.From)) {
		for _, chatUser := range(c.ChatUsersByLogin) {
			if(chatUser.IsOnline()) {
				reply.Messages = append(reply.Messages, chatUser.GetDetailsAsString())
			}
		}
	} else {
		reply.Message = "ERROR"
	}

	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) GetAllMyGroups(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	chatterSendable, found := c.SendToMap[query.From]
	chatUser, isChatUser := chatterSendable.(*ChatUser)
	if(found && isChatUser) {
		for _, chatGroup := range(chatUser.ChatGroups) {
			reply.Messages = append(reply.Messages, chatGroup.GetDetailsAsString())
		}
	} else {
		reply.Message = "ERROR"
	}
	
	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) GetAllMyActiveConnections(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	chatterSendable, found := c.SendToMap[query.From]
	chatUser, isChatUser := chatterSendable.(*ChatUser)
	if(found && isChatUser) {
		for _, chatConnection := range(chatUser.ChatConnections) {
			reply.Messages = append(reply.Messages, chatConnection.GetDetailsAsString())
		}
	} else {
		reply.Message = "ERROR"
	}
	
	c.SendById_unsafe(reply)
	c.Lock.Unlock()
}

func (c *ChatManager) SearchChatterBox(query ChatMessage) {
	reply := StdReplyToOrigin(query)
	c.Lock.Lock()

	if(!IsChatUserId(query.From)) {
		reply.Message = "ERROR"
		c.SendById_unsafe(reply)
	} else {
		if(IsChatId(query.Message)) {
			chatterSendable, found := c.SendToMap[query.Message]
			if(found) {
				reply.Messages = append(reply.Messages, chatterSendable.GetDetailsAsString())
			}
		} else {
			chatterBoxesByName, found := c.UsersAndGroups[query.Message]
			if(found) {
				for _, chatterBox := range chatterBoxesByName {
					reply.Messages = append(reply.Messages, chatterBox.GetDetailsAsString())
				}
			}
		}

		c.SendById_unsafe(reply)
	}

	c.Lock.Unlock()
}