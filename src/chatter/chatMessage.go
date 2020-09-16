package chatter

import (
	"time"
	"golang.org/x/net/websocket"
)

type ChatMessage struct {
	From string 		// id of the sender (can be a chat connection or a chat user)
	To string 			// id of the receiver (can be a chat connection or a chat user or a chat group)
	SentAt time.Time 	// when was the message sent (time stamp)
	Message string 		// string message content
}

func EmptyMessage() ChatMessage {
	return ChatMessage{}
}

func (c *ChatMessage) IsValid() bool {
	if( (IsChatConnectionId(c.From) || IsChatUserId(c.From)) && 
	(IsChatId(c.To) || IsChatConnectionId(c.To) || IsChatGroupId(c.To)) ) {
		return true
	}
	return false
}

func (c *ChatMessage) IsValidServerRequest() bool {
	if
}

var ChatMessageCodec = websocket.JSON;