package chatter

type ChatGroup struct {
	// name of the chat group
	Name string

	// add a message here to send it to every one else in the group
	InputMessage chan ChatMessage

	// any message received in the InputMessage is sent to every one except the sender
	ChatUsers []*ChatUsers

	// last message, received or sent or pinged
	LastMessage time.Time
}

func NewChatGroup(name string, users []*ChatUser) *ChatGroup {
	grp := &ChatGroup{
		Name:name,
		InputMessage:make(chan ChatMessage, 10),
		ChatUsers: users,
		LastMessage:time.Now(),
	}
	go grp.LoopOverChannelToPassMessagesToThisUser()
	return grp
}

func (grp *ChatGroup) SendMessage(msg ChatMessage) {
	grp.InputMessage <- msg
}

func (grp *ChatGroup) LoopOverChannelToPassMessages() {
	for msg := range grp.InputMessage {
		if(msg.To == grp.Name) {
			for user := range grp.ChatUser {
				if(msg.From != user.Name) {
					user.SendMessage(msg)
				}
			}
			grp.LastMessage = time.Now()
		}
	}
}