package chatter

type ChatterBox interface {
	GetId() string
	GetName() string
	SetName(name string)
	SendMessage(msg ChatMessage)
	LoopOverChannelToPassMessages()
	Destroy()
}

type ChatterBoxIndentity struct {
	Id string
	Name string
}

func (c *ChatterBoxIndentity) GetId() string {
	return c.Id
}

func (c *ChatterBoxIndentity) GetName() string {
	return c.Name
}

func (c *ChatterBoxIndentity) SetName(name string) {
	c.Name = name
}


