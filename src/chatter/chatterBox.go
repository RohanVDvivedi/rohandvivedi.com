package chatter

// A Chatter sendable exists only inside chat manager
type ChatterSendable interface {
	GetId() string
	SendMessage(msg ChatMessage)
	Destroy()
}

type ChatterBox interface {
	ChatterSendable
	GetName() string
	SetName(name string)
}

type Id struct {
	Id string
}

func (i *Id) GetId() string {
	return i.Id
}

type Name struct {
	Name string
}

func (n *Name) GetName() string {
	return n.Name
}

func (n *Name) SetName(name string) {
	n.Name = name
}


