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

	var chatConnection *ChatConnection = nil

	chatConnectionIntr, found := session.GlobalSessionStore.GetExistingSession(conn.Request()).GetValue("chat_connection")
	if(found) {
		chatConnectionTemp, isChatConnectionType := chatConnectionIntr.(*ChatConnection)
		if(isChatConnectionType) {
			chatConnection = chatConnectionTemp
		}
	}
	if(chatConnection == nil) {
		chatConnection = NewChatConnection();
		session.GlobalSessionStore.GetExistingSession(conn.Request()).SetValue("chat_connection", chatConnection)
	}

	chatConnection.Start(conn);

	for (true) {
		msg, err := chatUser.ReceiveMessage()
		if(err != nil) {
			chatConnection.Stop();
			break
		}
		if(msg.IsValidChatMessage()) {
			receiverUser := Chatters.GetChatterBoxById(msg)
		} else if (msg.IsValidServerRequest()) {
			Chatters.ServerMessagesToBeProcessed.Push(msg)
		}
	}

	chatConnection.WaitForShutdown();
	chatConnection.Stop();

	Chatters.DeleteChatterer(chatConnection.GetId());
}