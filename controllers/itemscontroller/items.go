package itemscontroller

import (
	"encoding/json"
	"net/http"
)

// Item is a model struct that contains items to be ordered by a customer
type Item = struct {
	ID    int64
	Name  string
	Price string
}

// GetItems gets all items in the database
func GetItems(w http.ResponseWriter, r *http.Request) {
	var items = []Item{}
	json.NewEncoder(w).Encode(items)
}
