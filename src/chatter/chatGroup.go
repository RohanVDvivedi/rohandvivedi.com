package chatter

type ChatGroup struct {
	Id
	Name

	ChatUsers map[string]*ChatUser
}

func NewChatGroup(name string) *ChatGroup {
	grp := &ChatGroup{
		Id: Id{GetNewChatUserId()},
		Name: Name{name},
		ChatUsers: make(map[string]*ChatUser),
	}
	return grp
}

func (grp *ChatGroup) SendMessage(msg ChatMessage) error {
	if(msg.To == grp.GetId()) {
		for _, user := range grp.ChatUsers {
			user.SendMessage(msg)
		}
	}
	return nil
}

func (grp *ChatGroup) Destroy() {
}

/* Joinery methods */
func (grp *ChatGroup) AddChatUser(c *ChatUser) {
	grp.ChatUsers[c.GetId()] = c
}
func (grp *ChatGroup) RemoveChatUser(c *ChatUser) {
	delete(grp.ChatUsers, c.GetId())
}