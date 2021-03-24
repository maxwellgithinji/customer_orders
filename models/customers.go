package models

import "time"

// Customer is a model struct that contains customer details
type Customer = struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
