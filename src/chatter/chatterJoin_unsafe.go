package chatter

// all the methods below are chatMAnager unsafe and must be called only while holding the lock
// all these functions return bool, which represent whether 

/* ChatUser and ChatGroup */
func isUserJoinedToGroup(cu *ChatUser, cg *ChatGroup) {
	return cu.HasChatGroup(cg) && cg.HasChatUser(cu)
}

func JoinUserToGroup(cu *ChatUser, cg *ChatGroup) {
	if(!isUserJoinedToGroup(cu, cg)) {
		cu.AddChatGroup(cg)
		cg.AddChatUser(cu)
		return true
	}
	return false
}

func BreakUserFromGroup(cu *ChatUser, cg *ChatGroup) {
	if(isUserJoinedToGroup(cu, cg)) {
		cu.RemoveChatGroup(cg)
		cg.RemoveChatUser(cu)
		return true
	}
	return false
}

/* ChatConnection and ChatUser */
func isConnectionJoinedToUser(cc *ChatConnection, cu *ChatUser) bool {
	return (cc.GetChatUser() == cu) && cu.HasChatConnection(cc)
}

func JoinConnectionToUser(cc *ChatConnection, cu *ChatUser) bool {
	if(!isConnectionJoinedToUser(cc, cu)) {
		cc.SetChatUser(cu)
		cu.AddChatConnection(cc)
		return true
	}
	return false
}

func BreakConnectionFromUser(cc *ChatConnection, cu *ChatUser) bool {
	if(isConnectionJoinedToUser(cc, cu)) {
		cc.RemoveChatUser(cu)
		cu.RemoveChatConnection(cc)
		return true
	}
	return false
}