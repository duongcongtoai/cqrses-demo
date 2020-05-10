package es

import (
	"fmt"
	"time"
)

type Event interface {
	Type() string
	Scan(interface{}) error
}
type EventMessage struct {
	AggregateID string    `json:"agg_id" db:"agg_id"`
	Type        string    `json:"type" db:"type"`
	Version     int       `json:"version" db:"version"`
	Metadata    []byte    `json:"meta_data" db:"meta_data"`
	Data        Event     `json:"data" db:"data"`
	CreatedOn   time.Time `json:"created_on" db:"created_on"`
	// CreatedOn   JSONTime    `json:"created_on"`
}

func NewEventMessage(aggregateID string, event Event, version int, createdOn time.Time) EventMessage {
	return EventMessage{
		aggregateID,
		event.Type(),
		version,
		[]byte{},
		event,
		createdOn,
	}
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(time.Time(t).Format(time.RFC3339))
	return []byte(stamp), nil
}

func (t JSONTime) UnmarshalJSON(b []byte) error {
	parsedTime, err := time.Parse(time.RFC3339, string(b))
	if err != nil {
		return err
	}
	t = JSONTime(parsedTime)
	return nil
}
