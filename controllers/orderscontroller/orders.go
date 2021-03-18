package orderscontroller

import (
	"encoding/json"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/models"
)

// GetOrders gets all customer orders in the database
func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders = []models.Order{}
	json.NewEncoder(w).Encode(orders)
}
