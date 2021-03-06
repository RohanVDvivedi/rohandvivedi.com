package chatter

// unsafe versions fo utility functions to be called from inside of other queries, when they have locks
func (c *ChatManager) InsertChatterer_unsafe(chatterer ChatterSendable) {
	// insert to the main map allowing us to send messages
	c.SendToMap[chatterer.GetId()] = chatterer

	// insertions by name for groups and users for find by name stuff
	chatterBox, isChatterBox := chatterer.(ChatterBox)
	if(isChatterBox) {
		_, chatterBoxesPresent := c.UsersAndGroups[chatterBox.GetName()]
		if(!chatterBoxesPresent) {
			c.UsersAndGroups[chatterBox.GetName()] = make(map[string]ChatterBox)
		}
		c.UsersAndGroups[chatterBox.GetName()][chatterBox.GetId()] = chatterBox
	}

	// insertion by name and public key for authentication
	chatUser, isChatUser:= chatterer.(*ChatUser)
	if(isChatUser) {
		c.ChatUsersByLogin[chatUser.GetName() + chatUser.PublicKey] = chatUser
	}
}

func (c *ChatManager) SendById_unsafe(msg ChatMessage) bool {
	chatterSendable, found := c.SendToMap[msg.To]
	if(found) {
		chatterSendable.SendMessage(msg);
	}
	return found
}

func (c *ChatManager) NotifyOnlineUsers_unsafe(notif ChatMessage) {
	for _, chatUser := range(c.ChatUsersByLogin) {
		if(chatUser.IsOnline()) {
			notif.To = chatUser.GetId()
			chatUser.SendMessage(notif)
		}
	}
}

func (c *ChatManager) DeleteChatterer_unsafe(Id string) {
	chatterSendable, found := c.SendToMap[Id]
	if(found) {
		var modifedChatUser *ChatUser = nil
		chatConnection, isChatConnection := chatterSendable.(*ChatConnection)
		if(isChatConnection && chatConnection.HasChatUser()) {
			modifedChatUser = chatConnection.GetChatUser()
		}

		chatterSendable.Destroy();

		delete(c.SendToMap, Id);
		chatterBox, isChatterBox := chatterSendable.(ChatterBox)
		if(isChatterBox) {
			delete(c.UsersAndGroups[chatterBox.GetName()], chatterBox.GetId());
		}
		chatUser, isChatUser := chatterSendable.(*ChatUser)
		if(isChatUser) {
			delete(c.ChatUsersByLogin, chatUser.GetName() + chatUser.PublicKey);
		}

		if(modifedChatUser != nil) {
			c.NotifyOnlineUsers_unsafe(ChatMessage{From:"server-event-update",Message:modifedChatUser.GetDetailsAsString()})
		}
	}
}

// Utility functions to help other queries
func StdReplyToOrigin(msg ChatMessage) ChatMessage {
	return ChatMessage{From:msg.To, To:msg.OriginConnection, ContextId: msg.MessageId, Message: "", Messages: []string{}}
}

func StdReplyToSender(msg ChatMessage) ChatMessage {
	return ChatMessage{From:msg.To, To:msg.From, ContextId: msg.MessageId, Message: "", Messages: []string{}}
}