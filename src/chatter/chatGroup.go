package chatter

import(
	"sync"
)

type ChatGroup struct {
	sync.RWMutex

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

func (grp *ChatGroup) GetDetailsAsString() string {
	return grp.GetId() + "," + grp.GetName()
}

// the message must be from a chat user who is part of the group or is the server
func (grp *ChatGroup) SendMessage(msg ChatMessage) {
	_, senderIsGroupMember := grp.ChatUsers[msg.From]
	if((senderIsGroupMember || IsChatManagerId(msg.From)) && (msg.To == grp.GetId())) {
		for _, user := range grp.ChatUsers {
			user.SendMessage(msg)
		}
	}
}

func (grp *ChatGroup) Destroy() {
}

/* Joinery methods */
func (grp *ChatGroup) HasChatUser(c *ChatUser) bool {
	_, found := grp.ChatUsers[c.GetId()]
	return found
}
func (grp *ChatGroup) AddChatUser(c *ChatUser) {
	grp.ChatUsers[c.GetId()] = c
}
func (grp *ChatGroup) RemoveChatUser(c *ChatUser) {
	delete(grp.ChatUsers, c.GetId())
}
func (grp *ChatGroup) GetChatUsersCount() int {
	return len(grp.ChatUsers)
}