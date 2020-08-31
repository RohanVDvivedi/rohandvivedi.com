package chatter

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"time"
	"fmt"
)

import (
	"rohandvivedi.com/src/session"
)

// map of all the chat users 
// from their name to the chatUser struct pointer
var map[string]*chatUser = {}
 
func ChatHandler(conn *websocket.Conn) {

	r := conn.Request()
	// the user must share his Name and a PublicKey
	
	r.

	fmt.Println(*s)
	
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