package models

// Order is a model struct of items ordered by a customer
type Order = struct {
	ID         int64   `json:"id"`
	CustomerID int64   `json:"customer_id"`
	ItemID     int64   `json:"item_id"`
	TotalPrice float64 `json:"total_price"`
	OderDate   string  `json:"order_date"`
}
