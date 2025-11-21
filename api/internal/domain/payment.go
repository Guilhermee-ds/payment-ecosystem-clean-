package domain

import "time"

type Payment struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	CardHash  string    `json:"card_hash"`
	CreatedAt time.Time `json:"created_at"`
}
