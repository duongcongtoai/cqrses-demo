package es

//CommandMessage Message used to send to eventstore implement this interface
type CommandMessage interface {
	ID() string
	Type() string
	Content() interface{}
}

//DescriptiveCommandMessage default implementation of commandmessage
type DescriptiveCommandMessage struct {
	command interface{}
	uuid    string
}

//Content ... simply return original command
func (d *DescriptiveCommandMessage) Content() interface{} {
	return d.command
}

func (d *DescriptiveCommandMessage) ID() string {
	return d.uuid
}

//Type ... return the type of the command for dispatching
func (d *DescriptiveCommandMessage) Type() string {
	return typeOf(d.command)
}

//NewCommandMessage return new implementation of descriptiveCommandmessage
func NewCommandMessage(uuid string, command interface{}) *DescriptiveCommandMessage {
	return &DescriptiveCommandMessage{
		command,
		uuid,
	}
}
