package chatter

import (
	"time"
)

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