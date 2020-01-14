package socket

import (
	"golang.org/x/net/websocket"
	"time"
	"fmt"
)

type resp struct {
	Message string
}
 
func Handler(conn *websocket.Conn) {
	
	go readFromConn(conn);

	for i := 0; i < 20; i++ {
		websocket.JSON.Send(conn, struct {Time time.Time; Iterator int}{Time: time.Now(), Iterator: i});

		// loop every 2 seconds
		time.Sleep(2 * time.Second);
	}

	conn.Close();
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