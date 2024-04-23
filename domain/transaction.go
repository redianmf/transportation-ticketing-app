package domain

import "time"

type Transaction struct {
	ID                   int       `json:"id"`
	UserId               int       `json:"user_id"`
	TransportationModeId int       `json:"transportation_mode_id"`
	TransactionReference string    `json:"transaction_reference"`
	Type                 string    `json:"type"`
	Status               string    `json:"status"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
