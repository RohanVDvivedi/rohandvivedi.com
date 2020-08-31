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

// a ping message : shoudl be sent by the client at regular interval to keep the socket connection alive
/* ping message must have all the fields empty
{"From":"","To":"","SentAt":time.Now,"Message":""}
*/
// this is the message that goes to and from a user to another
type chatMessage struct {
	from string

	to string

	sentAt time.Time

	message string
}