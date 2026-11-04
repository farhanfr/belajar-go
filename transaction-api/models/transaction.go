package models

type Transaction struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Amount int    `json:"amount"`
}
