package aggregates

import (
	"cqrses/registration/valueobject"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateOrder(t *testing.T) {
	Convey("Given an order id, conferenceid and seats info", t, func() {
		id := "orderId"
		conferenceID := "conferenceID"
		seatsInfo := valueobject.OrderSeatsInfo{
			"vip": {"vip", 3},
		}
		Convey("When create an order", func() {
			o := NewOrder()
			err := o.Create(id, conferenceID, seatsInfo)
			So(err, ShouldEqual, nil)
			// fmt.Println("here")
			Convey(fmt.Sprintf("Order id should be %s", id), func() {
				So(o.ID, ShouldEqual, id)
			})
			Convey(fmt.Sprintf("Order conference id should be %s", conferenceID), func() {
				So(o.ConferenceID, ShouldEqual, conferenceID)
			})
			Convey(fmt.Sprintf("Order seats items should be %v", seatsInfo), func() {
				So(o.SeatsInfo, ShouldEqual, seatsInfo)
			})
		})
	})
}

func TestAddSeat(t *testing.T) {
	Convey("Given an available order with 3 vip seats", t, func() {
		id := "orderId"
		conferenceID := "conferenceID"
		seatsInfo := valueobject.OrderSeatsInfo{
			"vip": {"vip", 3},
		}
		o := NewOrder()
		o.Create(id, conferenceID, seatsInfo)

	})
}
