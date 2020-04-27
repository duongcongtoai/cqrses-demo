package es

import (
	"fmt"
	"time"
)

type AggregateRoot interface {
	GetID() string
	PersistedVersion() int
	CurrentVersion() int
	// IncrementVersion()
	// TrackChange(Event)
	GetChanges() []EventMessage
	// ClearChanges()
	Apply(Event) error
	// Hydrate(Event) error
	TrackChange(Event)
}

type AggregateBase struct {
	ID      string
	version int
	changes []EventMessage
}

func AddEvent(a AggregateRoot, e Event) error {
	a.TrackChange(e)
	return a.Apply(e)
}

type EventApplyMethodNotFound string

func (e EventApplyMethodNotFound) Error() string {
	return fmt.Sprintf("Apply method for event %s not found", e)
}

func NewAggregateBase(id string) *AggregateBase {
	return &AggregateBase{
		ID:      id,
		changes: []EventMessage{},
		version: -1,
	}
}
func (a *AggregateBase) TrackChange(ev Event) {
	a.changes = append(a.changes, NewEventMessage(a.ID, ev, a.CurrentVersion()+1, time.Now()))
}

func (a *AggregateBase) GetID() string {
	return a.ID
}

func (a *AggregateBase) PersistedVersion() int {
	return a.version
}

func (a *AggregateBase) CurrentVersion() int {
	return a.version + len(a.changes)
}

// func (a *AggregateBase) IncrementVersion() {
// 	a.version++
// }

func (a *AggregateBase) GetChanges() []EventMessage {
	return a.changes
}

// func (a *AggregateBase) ClearChanges() {
// 	a.changes = []EventMessage{}
// }

// func (a *AggregateBase) Apply(e Event) error {
// 	if handler, ok := a.applymethods[e.Type()]; ok {
// 		a.TrackChange(e)
// 		handler(a, e)
// 		return nil
// 	}
// 	return EventApplyMethodNotFound(e.Type())
// }

// func (a *AggregateBase) Hydrate(e Event) error {
// 	return hydrate(a, e)
// }
