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
		msg, err := chatUser.ReceiveMessage()
		// if there happens to be any error in receiving a message, user created or not, we close the connection up
		if(err != nil) {
			break;
		}
		receiverUser, found := Chatters[msg.To]
		if(found) {
			receiverUser.SendMessage(msg)
		}
	}

	// remove 
	session.GlobalSessionStore.GetExistingSession(conn.Request()).RemoveValue("name");

	delete(Chatters, name);
}