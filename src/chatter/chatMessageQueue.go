package chatter

type ChatMessageQueue struct {

}

func NewChatMessageQueue() *ChatMessageQueue {
	return &ChatMessageQueue{
	}
}

func (c *ChatMessageQueue) Push(m ChatMessage) {

}

func (c *ChatMessageQueue) Pop() {
	
}

func (c *ChatMessageQueue) Top() ChatMessage {
	return ChatMessage{}
}

func (c *ChatMessageQueue) Destroy() {
	
}