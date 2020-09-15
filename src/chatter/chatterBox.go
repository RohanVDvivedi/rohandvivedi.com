package chatter

type ChatterBox interface {
    SendMessage(msg ChatMessage)
    LoopOverChannelToPassMessages()
}