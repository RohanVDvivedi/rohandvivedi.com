package chatter

import (
	"golang.org/x/net/websocket"
	"time"
	"sync"
)

type ChatConnection struct {
	MessagesToBeSent chan ChatMessage
	Connection *websocket.Conn
	ConnectionCloseWait sync.WaitGroup
	LastMessage time.Time
}

func NewChatConnection() *ChatConnection {
	return &ChatConnection{
		MessagesToBeSent: make(chan ChatMessage, 10),
		LastMessage:time.Now(),
	}
}

func (cconn *ChatConnection) Start(Connection *websocket.Conn, ReceivedMessages chan ChatMessage) {
	cconn.Connection = Connection
	cconn.ConnectionCloseWait.Add(2)
	go cconn.LoopOverChannelToSendMessages()
	go cconn.LoopOverToReceiveMessages(ReceivedMessages)
}

func (cconn *ChatConnection) SendMessage(msg ChatMessage) {
	cconn.MessagesToBeSent <- msg
}

func (cconn *ChatConnection) LoopOverToReceiveMessages(ReceivedMessages chan ChatMessage) {
	for(true) {
		msg := ChatMessage{}
		err := ChatMessageCodec.Receive(cconn.Connection, &msg)
		if(err != nil) {	// this could mean, connection closed or malformed chatMessage packet
			break
		}
		cconn.LastMessage = time.Now()
		ReceivedMessages <- msg
	}
	cconn.ConnectionCloseWait.Done()
	cconn.Connection.Close()
}

func (cconn *ChatConnection) LoopOverChannelToSendMessages() {
	for msg := range cconn.MessagesToBeSent {
		err := ChatMessageCodec.Send(cconn.Connection, msg)
		if(err != nil) { // this could mean, connection closed or lost
			break
		}
		cconn.LastMessage = time.Now()
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
}

func (cconn *ChatConnection) Destroy() {
	close(cconn.MessagesToBeSent)
}