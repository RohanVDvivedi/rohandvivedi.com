package chatter

import (
	"time"
	"golang.org/x/net/websocket"
)

type ChatMessage struct {
	// this is set by a connection when ever a message is created
	// the connection must set here its id
	OriginConnection string

	From string 		// id of the sender (can be a chat connection or a chat user)
	To string 			// id of the receiver (can be a chat connection or a chat user or a chat group)
	SentAt time.Time 	// when was the message sent (time stamp)
	Message string 		// string message content
}

func EmptyMessage() ChatMessage {
	return ChatMessage{}
}

func (c *ChatMessage) IsValidChatMessage() bool {
	if( (IsChatConnectionId(c.From) && IsChatConnectionId(c.To)) 	|| 
		(IsChatUserId(c.From) && IsChatUserId(c.To)) 				|| 
		(IsChatUserId(c.From) && IsChatGroupId(c.To)) 					) {
		return true
	}
	return false
}

func (c *ChatMessage) IsValidServerRequest() bool {
	if( (IsChatConnectionId(c.From) || IsChatUserId(c.From)) && (IsChatManagerId(c.To)) ) {
		return true
	}
	return false
}

var ChatMessageCodec = websocket.JSON;