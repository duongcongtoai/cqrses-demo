package aggregates

import (
	es "cqrses/es"
	orderevent "cqrses/registration/event"
	"cqrses/registration/valueobject"
	"fmt"
	"reflect"
)

type applyMethodsAggregateInstanceErr struct {
	expected string
	passedIn string
	method   string
}

func (a applyMethodsAggregateInstanceErr) Error() string {
	return fmt.Sprintf("Expected %s, %s passed in %s", a.expected, a.passedIn, a.method)
}

func NewOrder() *Order {
	return &Order{AggregateBase: &es.AggregateBase{}}
}

type Order struct {
	*es.AggregateBase
	ConferenceID string
	IsConfirmed bool
	OrderSeats []valueobject.SeatsQuantity
}

type OrderItems struct {
	Quantity uint
	SeatType string
}


func convert(s []OrderSeats) []valueobject.SeatsQuantity  {
	result := make(valueobject.SeatsQuantity, len(s))
	for index, item := range s {
		result[index] = orderevent.SeatsQuantity{item.Type, item.Quantity}
	}
	return result
}

func (o *Order) Create(id, conferenceID string, seats []OrderItems) error {
	seatsOrdered := convert(seats)
	return es.AddEvent(o, orderevent.OrderPlaced{
		id,
		seatsOrdered,
		seats,
	})
}

func (o *Order) UpdateSeats(seats []OrderItems) error {
	updatedSeats := convert(seats)
	return es.AddEvent(o, orderevent.SeatsUpdated{updatedSeats})
}

func (o *Order) MarkAsReserved(exp time.Time, reservedSeats []valueobject.SeatsQuantity)
func (o *Order) applyOrderPlaced(e orderevent.OrderPlaced) error {
	o.ID = e.ID
	o.ConferenceID = e.ConferenceID
	o.OrderSeats = e.Seats
	return nil
}

func (o *Order) applySeatsUpdated(e orderevent.SeatsUpdated) error {
	o.OrderSeats =
}

//Apply Implement Apply method
func (o *Order) Apply(e es.Event) error {
	switch e.(type) {
	case orderevent.OrderPlaced:
		ev := e.(orderevent.OrderPlaced)
		return o.applyOrderPlaced(ev)
	case orderevent.SeatsUpdated:
		ev := e.(orderevent.SeatsUpdated)
		return o.applySeatsUpdated(ev)
	default:
		return fmt.Errorf("Apply method for event %s not found", reflect.TypeOf(e))
	}
}
