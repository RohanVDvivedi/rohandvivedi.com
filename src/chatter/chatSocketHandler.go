package chatter

import (
	"golang.org/x/net/websocket"
)

import (
	//"rohandvivedi.com/src/session"
)

// map of all the chat users 
// from their name to the chatUser struct pointer
var Chatters = map[string]*ChatUser{}
 
func ChatHandler(conn *websocket.Conn) {

	defer conn.Close();

	r := conn.Request()
	// the user must share his Name

	name := ""
	nameList, existsName := r.URL.Query()["name"];
	if(!existsName) {
		return
	}
	name = nameList[0]

	_, chatUserSameNameExists := Chatters[name]
	if(chatUserSameNameExists) {
		return
	}

	chatUser := NewChatUser(name, conn)
	defer chatUser.DestroyChatUser()

	Chatters[name] = chatUser

	for (true) {
		msg := chatUser.ReceiveMessage()
		if(msg.From == chatUser.Name) {
			receiverUser, found := Chatters[msg.To]
			if(found) {
				receiverUser.SendMessage(msg)
			}
		}
	}

	delete(Chatters, name);

}