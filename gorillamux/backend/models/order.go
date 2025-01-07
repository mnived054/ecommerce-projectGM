package models

type Order struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	Products   string  `json:"products"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"created-at"`
}
