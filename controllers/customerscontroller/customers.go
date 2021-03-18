package customerscontroller

import (
	"fmt"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/auth"
	"github.com/maxwellgithinji/customer_orders/models"
	"github.com/maxwellgithinji/customer_orders/utils"
)

// GetCustomers gets all customers in the database
// @Summary Get all customers in the database
// @Description Get all customers in the database
// @Tags  Customers
// @Produce  json
// @Success 200 {object} []models.Customer{}
// @Router /auth/customers [get]
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers = []models.Customer{}
	utils.ResponseWithDataHelper(w, "200", "customers fetch successful", customers)
}

// Profile gets profile of currently logged in user
// @Summary Get profile gets profile of currently logged in user
// @Description Get profile gets profile of currently logged in user
// @Tags  Customers
// @Produce  json
// @Success 200 {object} models.Customer{}
// @Router /auth/profile [get]
func Profile(w http.ResponseWriter, r *http.Request) {
	session, err := auth.Store.Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	var customer = models.Customer{}
	fmt.Printf("%v", session.Values["profile"])
	// TODO: Combine customer details and profile details to build the customer response
	utils.ResponseWithDataHelper(w, "200", "customer fetch successful", customer)
}
