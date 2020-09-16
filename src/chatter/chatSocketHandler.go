package chatter

import (
	"golang.org/x/net/websocket"
)

import (
	"rohandvivedi.com/src/session"
)

// map of all the chat users 
// from their name to the chatUser struct pointer
var Chatters = NewChatManager()


// never call this functions outside
func ChatConnectionHandler(conn *websocket.Conn) {
	nameIntr, _ := session.GlobalSessionStore.GetExistingSession(conn.Request()).GetValue("name")
	name, _ := nameIntr.(string)

	chatConnection = NewChatConnection(conn);

	Chatters.InsertChatterer(chatConnection);

	session.GlobalSessionStore.GetExistingSession(conn.Request()).SetValue("chat_active", true)
	defer session.GlobalSessionStore.GetExistingSession(conn.Request()).SetValue("chat_active", false)

	for (true) {
		msg, err := chatConnection.ReceiveMessage()
		if(err != nil) {
			break
		}
		if(msg.IsValidChatMessage()) {
			receiverUser := Chatters.SendById(msg)
		} else if (msg.IsValidServerRequest()) {
			Chatters.ServerMessagesToBeProcessed.Push(msg)
		}
	}

	Chatters.DeleteChatterer(chatConnection.GetId());
}