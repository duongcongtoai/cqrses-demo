package es

type EventBus interface {
	PublishEvent(EventMessage)
	AddHandler(EventHandler, ...interface{})
}

type InternalEventBus struct {
	handlers map[string]map[EventHandler]struct{}
}

func NewInternalEventBus() *InternalEventBus {
	return &InternalEventBus{
		handlers: make(map[string]map[EventHandler]struct{}),
	}
}

func (b *InternalEventBus) PublishEvent(event EventMessage) {
	if handlers, ok := b.handlers[event.Type]; ok {
		for handler := range handlers {
			handler.Handle(event)
		}
	}
}

func (b *InternalEventBus) AddHandler(handler EventHandler, events ...interface{}) {
	for _, event := range events {
		typeName := typeOf(event)
		if _, ok := b.handlers[typeName]; !ok {
			b.handlers[typeName] = make(map[EventHandler]struct{})
		}
		b.handlers[typeName][handler] = struct{}{}
	}
}
