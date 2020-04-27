package es

type EventHandler interface {
	Handle(EventMessage)
}
