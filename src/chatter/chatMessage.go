package chatter

import (
	"time"
	"golang.org/x/net/websocket"
)

type ChatMessage struct {
	// this is set by a connection when ever a message is created
	// the connection must set here its id
	OriginConnection string 	`json:",omitempty"`

	From string  		`json:",omitempty"`	// id of the sender (can be a chat connection or a chat user)
	To string 			`json:",omitempty"`	// id of the receiver (can be a chat connection or a chat user or a chat group)
	SentAt time.Time 	`json:",omitempty"`	// when was the message sent (time stamp)

	MessageId string 	`json:",omitempty"`	// string message id
	Message string 		`json:",omitempty"`	// string message content
	Messages []string 	`json:",omitempty"` // list of string messages content
	ContextId string 	`json:",omitempty"` // string message context id

	/* Things are getting a little shifty here */
	// MessageId is an attribute to be used by chat users to uniquely identify a message
	// ContextId is an attribute of a message to be used by receiver to let the sender know what is it the reply or response to
	// let us say if you are using acknowledgement for message being received or message being read
	// for those acknowledgement messages the Context id must be equal to the message id of the corresponding message
	// a reply to a message must carry a ContextId equal to the MessageId of the message that is being replied to
	// once a server request is processed (log in or log out), the ContextId of the server response must equal to the MessageId of the server request
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