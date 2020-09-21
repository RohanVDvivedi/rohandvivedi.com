package chatter

// all the methods below are chatMAnager unsafe and must be called only while holding the lock
// all these functions return bool, which represent whether the operation was completed or the result
// the parameters must not be null

/* ChatConnection and ChatUser */
func isConnectionJoinedToUser(cc *ChatConnection, cu *ChatUser) bool {
	return (cc.GetChatUser() == cu) && cu.HasChatConnection(cc)
}

func JoinConnectionToUser(cc *ChatConnection, cu *ChatUser) bool {
	if(!isConnectionJoinedToUser(cc, cu) && cc.GetChatUser() == nil) {
		cc.SetChatUser(cu)
		cu.AddChatConnection(cc)
		return true
	}
	return false
}

func BreakConnectionFromUser(cc *ChatConnection, cu *ChatUser) bool {
	if(isConnectionJoinedToUser(cc, cu)) {
		cc.RemoveChatUser()
		cu.RemoveChatConnection(cc)
		return true
	}
	return false
}