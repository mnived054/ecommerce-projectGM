package models

type Cart struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
