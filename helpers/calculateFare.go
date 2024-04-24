package helpers

import (
	"math"

	"github.com/redianmf/transportation-ticketing-app/domain"
)

func CalculateFare(transport domain.TransportationModes, startPoint int, endPoint int) (amount int) {
	if transport.PriceCalculation == "FLAT" {
		amount = transport.BasePrice
	} else {
		pointPassed := int(math.Abs(float64(endPoint) - float64(startPoint)))
		amount = transport.BasePrice + (transport.AdditionalPrice * pointPassed)
	}

	return
}
