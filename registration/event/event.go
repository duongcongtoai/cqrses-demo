package event

import (
	"cqrses/registration/valueobject"
)

type OrderPlaced struct {
	ID           string
	ConferenceID string
	Seats        []valueobject.SeatsQuantity
}

type SeatsUpdated struct {
	Seats []valueobject.SeatsQuantity
}

func (s SeatsUpdated) Type() string {
	return "SeatsUpdated"
}

func (o OrderPlaced) Type() string {
	return "OrderPlaced"
}
