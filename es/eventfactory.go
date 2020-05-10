package es

import "fmt"

type EventFactory interface {
	GetEvent(string) Event
}

type DelegateEventFactory struct {
	eventFactories map[string]Event
}

func NewDelegateEventFactory() *DelegateEventFactory {
	return &DelegateEventFactory{
		eventFactories: make(map[string]Event),
	}
}

func (f *DelegateEventFactory) RegisterDelegate(delegate Event) error {
	if _, ok := f.eventFactories[delegate.Type()]; ok {
		return fmt.Errorf("Delegate for type %s has already been registered", delegate.Type())
	}

	f.eventFactories[delegate.Type()] = delegate
	return nil
}

func (f *DelegateEventFactory) GetEvent(typeName string) Event {
	if f, ok := f.eventFactories[typeName]; ok {
		return f
	}
	return nil
}
