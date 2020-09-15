package chatter

type ChatterBox interface {
	GetId() string
	GetName() string
	SendMessage(msg ChatMessage)
	LoopOverChannelToPassMessages()
	Destroy()
}
