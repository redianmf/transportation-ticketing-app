package domain

import "time"

type Wallet struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
