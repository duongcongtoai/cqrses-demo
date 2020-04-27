package registration

import (
	"cqrses/registration/valueobject"

	"github.com/google/uuid"
)

type PricingService interface {
	CalculateTotal(conferenceId uuid.UUID, seats valueobject.SeatsQuantity) OrderTotal
}

type OrderTotal struct {
	Lines      []SeatsOrderLine
	TotalPrice float64
}
type SeatsOrderLine struct {
	SeatType  uuid.UUID
	UnitPrice float64
	Quantity  uint
}

type DefaultPricingService struct {
}

func (s DefaultPricingService) CalculateTotal(confernceId uuid.UUID, eats valueobject.SeatsQuantity) OrderTotal {

}
