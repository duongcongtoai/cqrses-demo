package es

type CommandHandler interface {
	Handle(CommandMessage) error
}
