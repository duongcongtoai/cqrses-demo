package es

import "fmt"

type AggregateFactory interface {
	GetAggregate(string, string) AggregateRoot
}

type DelegateAggregateFactory struct {
	delegates map[string]func(string) AggregateRoot
}

func NewDelegateAggregateFactory() *DelegateAggregateFactory {
	return &DelegateAggregateFactory{
		delegates: make(map[string]func(string) AggregateRoot),
	}
}

func (d *DelegateAggregateFactory) RegisterDelegate(aggType string, delegate func(string) AggregateRoot) error {
	if _, ok := d.delegates[aggType]; ok {
		return fmt.Errorf("Delegate for this type already registered: %s", aggType)
	}

	d.delegates[aggType] = delegate
	return nil
}

func (d *DelegateAggregateFactory) GetAggregate(typeName string, id string) AggregateRoot {
	if f, ok := d.delegates[typeName]; ok {
		return f(id)
	}
	return nil
}
