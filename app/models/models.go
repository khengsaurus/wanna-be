package models

import (
	"time"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Expense struct {
	Id       int       `json:"id"`
	UserId   int       `json:"userId"`
	Currency string    `json:"currency"`
	Amount   float32   `json:"amount"`
	Note     string    `json:"note"`
	Date     time.Time `json:"date"`
}
