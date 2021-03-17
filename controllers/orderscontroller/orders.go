package orderscontroller

import (
	"encoding/json"
	"net/http"
)

// Order is a model struct of items ordered by a customer
type Order = struct {
	ID         int64
	CustomerID int64
	ItemID     int64
	Quantity   int64
	Amount     float64
	OderDate   string
}

// GetOrders gets all customer orders in the database
func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders = []Order{}
	json.NewEncoder(w).Encode(orders)
}
