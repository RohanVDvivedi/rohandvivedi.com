package chatter

import (
	"golang.org/x/net/websocket"
	"time"
	"errors"
)

type ChatUser struct {
	Id string
	Name string
	PublicKey string

	MessagesToBeSent chan ChatMessage
	Connection *websocket.Conn
	LastMessage time.Time

	ChatGroups map[string]*ChatGroup
}

func NewChatUser(name string, publicKey string, connection *websocket.Conn) *ChatUser {
	user := &ChatUser{
		Id: GetNewChatUserId(),
		Name: name,
		MessagesToBeSent: make(chan ChatMessage, 10),
		PublicKey:publicKey,
		Connection:connection,
		LastMessage:time.Now(),
		ChatGroups:make(map[string]*ChatGroup),
	}
	go user.LoopOverChannelToPassMessages()
	return user
}

func (user *ChatUser) GetId() string {
	return user.Id
}

func (user *ChatUser) GetName() string {
	return user.Name
}

func (user *ChatUser) SendMessage(msg ChatMessage) {
	user.MessagesToBeSent <- msg
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

func (user *ChatUser) LoopOverChannelToPassMessages() {
	for msg := range user.MessagesToBeSent {
		if(msg.To == user.Name) {
			err := ChatMessageCodec.Send(user.Connection, msg)
			if(err != nil) { // this could mean, connection closed or lost
				user.Destroy()
				return
			}
			user.LastMessage = time.Now()
		}
	}
}

func (user *ChatUser) Destroy() {
	close(user.MessagesToBeSent)
	user.Connection.Close()
}