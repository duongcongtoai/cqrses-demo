package es

type Command interface {
	Type() string
}

type CommandBus interface {
	Dispatch(Command) error
}
