package chatter

func (c *ChatManager) GetAllUsers(query ChatMessage) ChatMessage {
	c.Lock.Lock()

	reply := StdReplyToOrigin(query)
	if(IsChatUserId(msg.From)) {
		for _, chatUser := range(c.ChatUsersMapped) {
			reply.Messages = append(reply.Messages, chatUser.GetId() + "," + chatUser.GetName() + "," + strconv.Itoa(chatUser.GetChatConnectionCount()))
		}
	} else {
		reply.Message = "ERROR"
	}

	c.Lock.Unlock()
	return reply
}

func (c *ChatManager) GetAllGroups(query ChatMessage) ChatMessage {
	reply := StdReplyToOrigin(query)
	if(IsChatUserId(msg.From)) {
		/*for _, chatUser := range(c.ChatUsersMapped) {
			reply.Messages = append(reply.Messages, chatUser.GetId() + "," + chatUser.GetName() + "," + strconv.Itoa(chatUser.GetChatConnectionCount()))
		}*/
	} else {
		reply.Message = "ERROR"
	}
	return reply
}

func (c *ChatManager) GetAllActiveUsers(query ChatMessage) ChatMessage {
	reply := StdReplyToOrigin(query)
	if(IsChatUserId(msg.From)) {
		for _, chatUser := range(c.ChatUsersMapped) {
			if(chatUser.GetChatConnectionCount() > 0) {
				reply.Messages = append(reply.Messages, chatUser.GetId() + "," + chatUser.GetName() + "," + strconv.Itoa(chatUser.GetChatConnectionCount()))
			}
		}
	} else {
		reply.Message = "ERROR"
	}
	return reply
}

func (c *ChatManager) GetAllMyGroups(query ChatMessage) ChatMessage {
	reply := StdReplyToOrigin(query)
	chatterSendable, found := c.Chatters[msg.From]
	chatUser, isChatUser := chatterSendable.(*ChatUser)
	if(found && isChatUser) {
		for _, chatGroup := range(chatUser.ChatGroups) {
			reply.Messages = append(reply.Messages, chatGroup.GetId() + "," + chatGroup.GetName())
		}
	} else {
		reply.Message = "ERROR"
	}
	return reply;
}

func (c *ChatManager) GetAllMyActiveConnections(query ChatMessage) ChatMessage {
	reply := StdReplyToOrigin(query)
	chatterSendable, found := c.Chatters[msg.From]
	chatUser, isChatUser := chatterSendable.(*ChatUser)
	if(found && isChatUser) {
		for _, chatConnection := range(chatUser.ChatConnections) {
			reply.Messages = append(reply.Messages, chatConnection.GetId())
		}
	} else {
		reply.Message = "ERROR"
	}
	return reply
}

func (c *ChatManager) SearchChatterBox(query ChatMessage) ChatMessage {
	reply := StdReplyToOrigin(query)
	chatterSendable, found := c.Chatters[msg.Message]
	chatterBox, isChatterBox := chatterSendable.(ChatterBox)
	if(found && isChatterBox) {
		reply.Message = chatterBox.GetName()
	} else {
		reply.Message = "ERROR"
	}
	return reply
}