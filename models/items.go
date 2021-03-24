package models

import "time"

// Item is a model struct that contains items to be ordered by a customer
type Item = struct {
	ID        int64     `json:"id"`
	Item      string    `json:"item"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// ItemPost is a model struct that contains items to be ordered by a customer
type ItemPost = struct {
	Item  string `json:"item"`
	Price int64  `json:"price"`
}
