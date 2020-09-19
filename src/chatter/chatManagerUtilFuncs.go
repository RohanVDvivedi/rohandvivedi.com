package chatter

// unsafe versions fo utility functions to be called from inside of other queries, when they have locks
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

func (c *ChatManager) SendById_unsafe(msg ChatMessage) bool {
	chatterSendable, found := c.Chatters[msg.To]
	if(found) {
		chatterSendable.SendMessage(msg);
	}
	return found
}

func (c *ChatManager) DeleteChatterer_unsafe(Id string) {
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
	}
}

// Utility functions to help other queries
func StdReplyToOrigin(ChatMessage msg) ChatMessage {
	return ChatMessage{From:msg.To, To:msg.OriginConnection, ContextId: msg.MessageId, Message: "", Messages: []string{}}
}

func StdReplyToSender(ChatMessage msg) ChatMessage {
	return ChatMessage{From:msg.To, To:msg.From, ContextId: msg.MessageId, Message: "", Messages: []string{}}
}

func GetDetailsString(cs ChatterSendable) string {
	return ""
}