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

	chatConnection := NewChatConnection(conn);
	defer chatConnection.Destroy()

	Chatters.InsertChatterer(chatConnection);
	defer Chatters.DeleteChatterer(chatConnection.GetId());

	session.GlobalSessionStore.GetExistingSession(conn.Request()).SetValue("chat_conn_active", true)
	defer session.GlobalSessionStore.GetExistingSession(conn.Request()).SetValue("chat_conn_active", false)

	name, publicKey, isAuthenticatable := chatConnection.GetNameAndPublicKey()
	if(isAuthenticatable) {
		Chatters.ServerMessagesToBeProcessed.Push(ChatMessage{
				OriginConnection:chatConnection.GetId(),
				From:chatConnection.GetId(),To:"server-login-as-chat-user",
				Message:name + "," + publicKey,
			})
	}

	for (true) {
		msg, err := chatConnection.ReceiveMessage()
		if(err != nil) {
			break
		}
		Chatters.AddChatMessageToChatManagersProcessingQueue(msg)
	}
}