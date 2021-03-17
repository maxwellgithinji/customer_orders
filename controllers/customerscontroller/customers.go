package customerscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/auth"
)

// Customer is a model struct that contains customer details
type Customer = struct {
	ID    int64
	Name  string
	Email string
	Code  string
}

// GetCustomers gets all customers in the database
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers = []Customer{}
	json.NewEncoder(w).Encode(customers)
}

// Profile gets profile of currently logged in user
func Profile(w http.ResponseWriter, r *http.Request) {
	// Initialize session
	err := auth.InitSession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := auth.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var customer = Customer{}
	fmt.Printf("%v", session.Values["profile"])
	// TODO: Combine customer details and profile details to build the customer response
	json.NewEncoder(w).Encode(customer)
}
