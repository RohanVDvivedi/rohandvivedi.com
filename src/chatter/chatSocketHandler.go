package chatter

import (
	"golang.org/x/net/websocket"
)

import (
	"rohandvivedi.com/src/session"
)

// map of all the chat users 
// from their name to the chatUser struct pointer
var Chatters = map[string]*ChatUser{}
 
func ChatHandler(conn *websocket.Conn) {
	defer conn.Close();

	// the user must have a Name

	nameIntr, _ := session.GlobalSessionStore.GetExistingSession(conn.Request()).GetValue("name")
	name, _ := nameIntr.(string)

	_, chatUserSameNameExists := Chatters[name]
	if(chatUserSameNameExists) {
		return
	}

	chatUser := NewChatUser(name, conn)
	defer chatUser.DestroyChatUser()

	Chatters[name] = chatUser

	for (true) {
		msg := chatUser.ReceiveMessage()
		receiverUser, found := Chatters[msg.To]
		if(found) {
			receiverUser.SendMessage(msg)
		}
	}

	delete(Chatters, name);

}