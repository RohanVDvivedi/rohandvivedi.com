package chatter

type ChatUser struct {
	Id
	Name
	PublicKey string

	MessagesPendingToBeSent *ChatMessageQueue

	ChatConnections map[string]*ChatConnection

	ChatGroups map[string]*ChatGroup
}

func NewChatUser(name string, publicKey string) *ChatUser {
	user := &ChatUser{
		Id: Id{GetNewChatUserId()},
		Name: Name{name},
		PublicKey: publicKey,
		MessagesPendingToBeSent: NewChatMessageQueue(),
		ChatGroups: make(map[string]*ChatGroup),
	}
	return user
}

func (user *ChatUser) SendMessage(msg ChatMessage) error {
	_, usersGroup := user.ChatGroups[msg.To]
	if(msg.To == user.GetId() || usersGroup) {
		sentTo := 0
		for _, cconn := range user.ChatConnections {
			err := cconn.SendMessage(msg)
			if(err != nil) {
				sentTo += 1
			} else {
				cconn.Destroy()
			}
		}
		if(sentTo == 0) {
			user.MessagesPendingToBeSent.Push(msg)
		}
	}
	return nil
}

func (user *ChatUser) Destroy() {
}