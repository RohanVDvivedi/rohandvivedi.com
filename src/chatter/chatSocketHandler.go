package chatter

import (
	"fmt"
	"golang.org/x/net/websocket"
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

	name, publicKey, isAuthenticatable := chatConnection.GetNameAndPublicKey()
	if(isAuthenticatable) {
		Chatters.ServerMessagesToBeProcessed.Push(ChatMessage{
				OriginConnection:chatConnection.GetId(),
				From:chatConnection.GetId(),To:"server-login-as-chat-user",
				Messages: []string{name, publicKey},
			})
	}

	for (true) {
		msg, err := chatConnection.ReceiveMessage()
		if(err != nil) {
			fmt.Println(err)
			break
		}
		Chatters.AddChatMessageToChatManagersProcessingQueue(msg)
	}
}