package es

import "fmt"

type AggregateFactory interface {
	GetAggregate(string, string) AggregateRoot
}

type DelegateAggregateFactory struct {
	delegates map[string]func(string) AggregateRoot
}

func NewDelegteAggregateFactory() *DelegateAggregateFactory {
	return &DelegateAggregateFactory{
		delegates: make(map[string]func(string) AggregateRoot),
	}
}

func (d *DelegateAggregateFactory) RegisterDelegate(agg AggregateRoot, delegate func(string) AggregateRoot) error {
	typeName := typeOf(agg)
	if _, ok := d.delegates[typeName]; ok {
		return fmt.Errorf("Delegate for this type already registered: %s", typeName)
	}

	d.delegates[typeName] = delegate
	return nil
}

func (d *DelegateAggregateFactory) GetAggregate(typeName string, id string) AggregateRoot {
	if f, ok := d.delegates[typeName]; ok {
		return f(id)
	}
	return nil
}
