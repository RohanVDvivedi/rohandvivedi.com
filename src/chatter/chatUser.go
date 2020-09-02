package chatter

import (
	"golang.org/x/net/websocket"
	"time"
	"errors"
)

type ChatUser struct {
	// name of the chat user client
	Name string

	// the active web socket connection to the user
	Connection *websocket.Conn

	// sender needs to write to this channel,
	// if the message is referred to: to the given user identified by the Name
	// the chat user's go rputing dedicated to the given user will read from the channel 
	// and writ the message to the connection
	InputMessage chan ChatMessage

	// last message, received or sent or pinged
	// this lets us identify id the user went idle, for a long time say 5 mins
	// if so we close the connection
	LastMessage time.Time
}

func NewChatUser(name string, connection *websocket.Conn) *ChatUser {
	user := &ChatUser{
		Name:name,
		Connection:connection,
		InputMessage:make(chan ChatMessage, 10),
		LastMessage:time.Now(),
	}
	go user.LoopOverChannelToPassMessagesToThisUser()
	return user
}

func (user *ChatUser) DestroyChatUser() {
	close(user.InputMessage)
	user.Connection.Close()
}

func (user *ChatUser) SendMessage(msg ChatMessage) {
	user.InputMessage <- msg
}

func (user *ChatUser) ReceiveMessage() (ChatMessage, error) {
	msg := ChatMessage{}
	err := ChatMessageCodec.Receive(user.Connection, &msg)
	if(err != nil) {	// this could mean, connection closed or malformed chatMessage packet
		return msg, err;
	}
	user.LastMessage = time.Now()
	if(msg.From != user.Name) { // user maliciously pretending to be someone else
		return ChatMessage{}, errors.New("Malicious user: wrong user name in from attribute")
	}
	return msg, nil
}

func (user *ChatUser) LoopOverChannelToPassMessagesToThisUser() {
	for msg := range user.InputMessage {
		if(msg.To == user.Name) {
			err := ChatMessageCodec.Send(user.Connection, msg)
			if(err != nil) { // this could mean, connection closed or lost
				user.DestroyChatUser()
				return
			}
			user.LastMessage = time.Now()
		}
	}
}