package domain

import "time"

type TransportationModes struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	BasePrice        int       `json:"base_price"`
	AdditionalPrice  int       `json:"additional_price"`
	PriceCalculation string    `json:"price_calculation"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
