package itemscontroller

import (
	"encoding/json"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/models"
)

// GetItems gets all items in the database
func GetItems(w http.ResponseWriter, r *http.Request) {
	var items = []models.Item{}
	json.NewEncoder(w).Encode(items)
}
