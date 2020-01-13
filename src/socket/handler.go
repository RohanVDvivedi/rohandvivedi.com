package socket

import (
	"golang.org/x/net/websocket"
	"encoding/json"
	"time"
	"fmt"
)

type resp struct {
	Message string
}
 
func Handler(conn *websocket.Conn) {
	
	go readFromConn(conn);

	for i := 0; i < 20; i++ {
		// json marshal with time
		sendbuff, err := json.Marshal(struct {Time time.Time; Iterator int}{Time: time.Now(), Iterator: i});

		// send serialized time
		if(err == nil) {
			conn.Write(sendbuff);
		} else {
			conn.Write([]byte("{}"));
		}

		// loop every 2 seconds
		time.Sleep(2 * time.Second);
	}

	conn.Close();
}

func readFromConn(conn *websocket.Conn) {
	res := resp{Message: ""};
	res_buff := []byte{};

	for {
		n, err := conn.Read(res_buff);
		if(err == nil) {
			json.Unmarshal(res_buff, &res);
			fmt.Printf("%d : %s\n", n, res.Message);
		} else {
			fmt.Println("Error reading response on socket");
		}
	}
}