package chatter

import (
	"time"
	"golang.org/x/net/websocket"
)

// this is the message that goes to and from a user to another
// or one user to a group
type ChatMessage struct {
	From string
	To string
	SentAt time.Time
	Message string
}

func EmptyMessage() ChatMessage {
	return ChatMessage{}
}

var ChatMessageCodec = websocket.JSON;