package chatter

import(
	"time"
)

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

	chatterer.SendMessage(ChatMessage{From:"server-chatters-creator",To:chatterer.GetId(),SentAt:time.Now()})
}

func (c *ChatManager) SendById_unsafe(msg ChatMessage) bool {
	chatterSendable, found := c.SendToMap[msg.To]
	if(found) {
		chatterSendable.SendMessage(msg);
	}
	return found
}

func (c *ChatManager) DeleteChatterer_unsafe(Id string) {
	chatterSendable, found := c.SendToMap[Id]
	if(found) {
		delete(c.SendToMap, Id);

		chatterBox, isChatterBox := chatterSendable.(ChatterBox)
		if(isChatterBox) {
			delete(c.UsersAndGroups[chatterBox.GetName()], chatterBox.GetId());
		}

		chatUser, isChatUser := chatterSendable.(*ChatUser)
		if(isChatUser) {
			delete(c.ChatUsersByLogin, chatUser.GetName() + chatUser.PublicKey);
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

func GetDetailsAsString(cs ChatterSendable) string {
	chatConnection, isChatConnection := cs.(ChatConnection)
	if(isChatConnection) {
		return chatConnection.GetId()
	}

	chatUser, isChatUser := cs.(ChatUser)
	if(isChatUser) {
		return chatUser.GetId() + "," + chatUser.GetName() + "," + strconv.Itoa(chatUser.GetChatConnectionCount()))
	}

	chatGroup, isChatGroup := cs.(ChatGroup)
	if(isChatGroup) {
		return chatGroup.GetId() + "," + chatGroup.GetName()
	}
}