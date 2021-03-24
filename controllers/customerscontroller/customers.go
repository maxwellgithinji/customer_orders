package customerscontroller

import (
	"fmt"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/utils"
)

// GetCustomers gets all customers in the database
// @Summary Get all customers in the database
// @Description Get all customers in the database
// @Tags  Customers
// @Produce  json
// @Success 200 {object} []models.Customer{}
// @Router /auth/customers [get]s
func (*customercontroller) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := customerService.FindAllCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	utils.ResponseWithDataHelper(w, "200", "customers fetch successful", customers)
}

// Profile gets profile of currently logged in user
// @Summary Get profile gets profile of currently logged in user
// @Description Get profile gets profile of currently logged in user
// @Tags  Customers
// @Produce  json
// @Success 200 {object} models.Customer{}
// @Router /auth/profile [get]
func (*customercontroller) Profile(w http.ResponseWriter, r *http.Request) {
	err := openIDAuthService.InitSession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := openIDAuthService.NewStore().Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile := session.Values["profile"]
	email := fmt.Sprintf("%v", profile.(map[string]interface{})["email"])
	customer, err := customerService.FindACustomerByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseWithDataHelper(w, "200", "customer fetch successful", customer)
}
