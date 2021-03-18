package models

// Order is a model struct of items ordered by a customer
type Order = struct {
	ID         int64
	CustomerID int64
	ItemID     int64
	Quantity   int64
	Amount     float64
	OderDate   string
}
