package models

type Profile struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
