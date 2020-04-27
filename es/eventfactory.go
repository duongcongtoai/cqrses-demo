package es

import "fmt"

type EventFactory interface {
	GetEvent(string) Event
}

type DelegateEventFactory struct {
	eventFactories map[string]func() Event
}

func NewDelegateEventFactory() *DelegateEventFactory {
	return &DelegateEventFactory{
		eventFactories: make(map[string]func() Event),
	}
}

func (f *DelegateEventFactory) RegisterDelegate(event Event, delegate func() Event) error {
	if _, ok := f.eventFactories[event.Type()]; ok {
		return fmt.Errorf("Delegate for type %s has already been registered", event.Type())
	}

	f.eventFactories[event.Type()] = delegate
	return nil
}

func (f *DelegateEventFactory) GetEvent(typeName string) Event {
	if f, ok := f.eventFactories[typeName]; ok {
		return f()
	}
	return nil
}
