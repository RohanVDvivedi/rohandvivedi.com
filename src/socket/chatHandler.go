package socket

import (
	"golang.org/x/net/websocket"
	"time"
	"fmt"
)


type ChatUser struct {

	// name of the chat user client
	Name string

	// the active web socket connection to the user
	Connection *websocket.Conn

	// sender needs to write to this channel,
	// if the message is referred to: to the given user identified by the Name
	InputMessage chan ChatMessage

	// last message, received or sent or pinged
	// this lets us identify id the user went idle, for a long time say 5 mins
	// if so we close the connection
	LastMessage time.Time
}


// a ping message : shoudl be sent by the client at regular interval to keep the socket connection alive
/* ping message must have all the fields empty
{"From":"","To":"","SentAt":time.Now,"Message":""}
*/
// this is the message that goes to and from a user to another
type ChatMessage struct {
	From string

	To string

	SentAt time.Time

	Message string
}
 
func ChatHandler(conn *websocket.Conn) {
	
	go readFromConn(conn);

	for i := 0; i < 5; i++ {
		websocket.JSON.Send(conn, struct {Time time.Time; Iterator int}{Time: time.Now(), Iterator: i});

		// loop every 2 seconds
		time.Sleep(1 * time.Second);
	}

	conn.Close();
}
type resp struct {
	Message string
}
func readFromConn(conn *websocket.Conn) {
	res := resp{Message: ""};
	for {
		err := websocket.JSON.Receive(conn, &res);
		if(err == nil) {
			fmt.Printf("%s\n", res.Message);
		} else {
			fmt.Println("Error reading response on socket");
			break;
		}
	}
}