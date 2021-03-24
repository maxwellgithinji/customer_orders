package models

// Item is a model struct that contains items to be ordered by a customer
type Item = struct {
	ID    int64  `json:"id"`
	Item  string `json:"item"`
	Price string `json:"price"`
}
