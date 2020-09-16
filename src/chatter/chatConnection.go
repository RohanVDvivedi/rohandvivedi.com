package chatter

import (
	"golang.org/x/net/websocket"
	"time"
	"sync"
	"errors"
)

type ChatConnection struct {
	MessagesToBeSent* ChatMessageQueue
	Connection *websocket.Conn
	ConnectionCloseWait sync.WaitGroup
	LastMessage time.Time
	IsActive bool
}

func NewChatConnection() *ChatConnection {
	return &ChatConnection{
		MessagesToBeSent: NewChatMessageQueue(),
		LastMessage:time.Now(),
		IsActive: false,
	}
}

func (cconn *ChatConnection) Start(Connection *websocket.Conn, ReceivedMessages chan ChatMessage) {
	cconn.IsActive = true
	cconn.Connection = Connection
	cconn.ConnectionCloseWait.Add(1)
	go cconn.LoopOverChannelToPassMessages()
}

func (cconn *ChatConnection) SendMessage(msg ChatMessage) {
	cconn.MessagesToBeSent.Push(msg)
}

func (cconn *ChatConnection) ReceiveMessage() (ChatMessage, error) {
	msg := ChatMessage{}
	if(!cconn.IsActive) {
		return msg, errors.New("Error connection not active")
	}
	err := ChatMessageCodec.Receive(cconn.Connection, &msg)
	if(err != nil) {	// this could mean, connection closed or malformed chatMessage packet
		return msg, err
	}
	cconn.LastMessage = time.Now()
	return msg, nil
}

func (cconn *ChatConnection) LoopOverToPassMessages() {
	for (true) {
		msg := cconn.MessagesToBeSent.Top()
		err := ChatMessageCodec.Send(cconn.Connection, msg)
		if(err != nil) { // this could mean, connection closed or lost
			break
		}
		cconn.LastMessage = time.Now()
		cconn.MessagesToBeSent.Pop()
	}
	cconn.ConnectionCloseWait.Done()
	cconn.Connection.Close()
}

func (cconn *ChatConnection) WaitForShutdown() {
	cconn.ConnectionCloseWait.Wait()
}

func (cconn *ChatConnection) Stop() {
	cconn.Connection.Close()
	cconn.WaitForShutdown()
	cconn.Connection = nil
	cconn.IsActive = false
}

func (cconn *ChatConnection) Destroy() {
	if(cconn.IsActive) {
		cconn.Stop()
	}
	cconn.MessagesToBeSent = nil
}