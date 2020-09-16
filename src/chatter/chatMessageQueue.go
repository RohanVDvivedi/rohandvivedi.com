package chatter

import (
	"sync"
)

type ChatMessageQueue struct {
	holderLock sync.Mutex
	Holder []ChatMessage
	holderEmptyWait *sync.Cond
}

func NewChatMessageQueue() *ChatMessageQueue {
	cmq := &ChatMessageQueue{Holder: make([]ChatMessage, 10)}
	cmq.holderEmptyWait = sync.NewCond(&(cmq.holderLock))
	return cmq
}

func (cmq *ChatMessageQueue) Push(m ChatMessage) {
	cmq.holderLock.Lock()
	cmq.Holder = append(cmq.Holder, m)
	cmq.holderEmptyWait.Signal()
	cmq.holderLock.Unlock()
}

func (cmq *ChatMessageQueue) Pop() {
	cmq.holderLock.Lock()
	for(len(cmq.Holder) == 0) {
		cmq.holderEmptyWait.Wait()
	}
	cmq.Holder[0] = EmptyMessage()
	cmq.Holder = cmq.Holder[1:]
	if(len(cmq.Holder) > 0) {
		cmq.holderEmptyWait.Signal()
	}
	cmq.holderLock.Unlock()
}

func (cmq *ChatMessageQueue) Top() ChatMessage {
	cmq.holderLock.Lock()
	for(len(cmq.Holder) == 0) {
		cmq.holderEmptyWait.Wait()
	}
	m := cmq.Holder[0]
	cmq.holderEmptyWait.Signal()
	cmq.holderLock.Unlock()
	return m
}