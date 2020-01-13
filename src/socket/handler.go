package socket

import (
	"golang.org/x/net/websocket"
	"encoding/json"
	"time"
)
 
func Handler(conn *websocket.Conn) {
	for i := 0; i < 20; i++ {
		// json marshal with time
		json, err := json.Marshal(struct {Time time.Time; Iterator int}{Time: time.Now(), Iterator: i});

		// send serialized time
		if(err == nil) {
			conn.Write(json);
		} else {
			conn.Write([]byte("{}"));
		}

		// loop every 2 seconds
		time.Sleep(2 * time.Second);
	}

	conn.Close();
}