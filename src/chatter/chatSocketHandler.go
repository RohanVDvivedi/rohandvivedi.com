package chatter

import (
	"golang.org/x/net/websocket"
)

import (
	"rohandvivedi.com/src/session"
)

// map of all the chat users 
// from their name to the chatUser struct pointer
var Chatters = Chatterers{Chatters:make(map[string]*ChatUser)}


// never call this functions outside
func ChatConnectionHandler(conn *websocket.Conn) {
	defer conn.Close();

	nameIntr, _ := session.GlobalSessionStore.GetExistingSession(conn.Request()).GetValue("name")
	name, _ := nameIntr.(string)

	chatUser := Chatters.InsertUniqueChatUserByName(name, conn)
	if(chatUser == nil) {
		return
	}
	defer Chatters.DeleteChatUserByName(name);

	session.GlobalSessionStore.GetExistingSession(conn.Request()).SetValue("chat_active", true)
	defer session.GlobalSessionStore.GetExistingSession(conn.Request()).SetValue("chat_active", false)

	for (true) {
		msg, err := chatUser.ReceiveMessage()
		// if there happens to be any error in receiving a message, user created or not, we close the connection up
		if(err != nil) {
			break;
		}
		receiverUser := Chatters.GetChatUserByName(msg.To)
		if(receiverUser != nil) {
			receiverUser.SendMessage(msg)
		}
	}
}