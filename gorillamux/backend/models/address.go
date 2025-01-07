package models

// Address represents a user's shipping address
type Address struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}
