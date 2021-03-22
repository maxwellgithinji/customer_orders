package models

// Item is a model struct that contains items to be ordered by a customer
type Item = struct {
	ID    int64
	Name  string
	Price string
}
