package chatter

import (
	"golang.org/x/net/websocket"
)

// map of all the chat users 
// from their name to the chatUser struct pointer
var Chatters = NewChatManager()


// never call this functions outside
func ChatConnectionHandler(conn *websocket.Conn) {
	chatConnection := NewChatConnection(conn);

	Chatters.InsertChatterer(chatConnection);
	defer Chatters.DeleteChatterer(chatConnection.GetId());

	name, publicKey, isAuthenticatable := chatConnection.GetNameAndPublicKey()
	if(isAuthenticatable) {
		// draft a login server message and send it on behalf of the client
		Chatters.ProcessMessage(ChatMessage{
			OriginConnection:chatConnection.GetId(),
			From:chatConnection.GetId(),To:"server-login-as-chat-user",
			Messages:[]string{name, publicKey},
		})
	}

	for (true) {
		msg, err := chatConnection.ReceiveMessage()
		if(err != nil) {
			break
		}
		Chatters.ProcessMessage(msg)
	}
}