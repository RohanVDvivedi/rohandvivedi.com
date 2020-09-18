package chatter

import (
	"sync"
)

type ChatMessageQueue struct {
	holderLock sync.Mutex
	Holder []ChatMessage

	// when the queue is empty, wait here
	// and you would be notified if anything is pushed to the queue
	holderEmptyWait *sync.Cond

	// set tot true of the queue must pause any consumption
	// push operation can still be issued to this queue
	PausedToPop bool
	// wait on this condition varibale if you are performing a pop and the queue has been popped
	queuePausedWait *sync.Cond
}

func NewChatMessageQueue() *ChatMessageQueue {
	cmq := &ChatMessageQueue{Holder: make([]ChatMessage, 10), PausedToPop: false}
	cmq.holderEmptyWait = sync.NewCond(&(cmq.holderLock))
	cmq.queuePausedWait = sync.NewCond(&(cmq.holderLock))
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

	// wait of queue holder is empty
	for(len(cmq.Holder) == 0 || cmq.PausedToPop) {
		if(len(cmq.Holder) == 0) {
			cmq.holderEmptyWait.Wait()
		} else if (cmq.PausedToPop) {
			cmq.queuePausedWait.Wait()
		}
	}

	cmq.Holder[0] = EmptyMessage()
	cmq.Holder = cmq.Holder[1:]
	if(len(cmq.Holder) > 0) {
		cmq.holderEmptyWait.Signal()
	}
	cmq.holderLock.Unlock()
}

func (cmq *ChatMessageQueue) PausePop() {
	cmq.holderLock.Lock()
	cmq.PausedToPop = true
	cmq.holderLock.Unlock()
}

func (cmq *ChatMessageQueue) ResumePop() {
	cmq.holderLock.Lock()
	cmq.PausedToPop = false
	cmq.queuePausedWait.Broadcast()
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