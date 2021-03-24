package customerscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/models"
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
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	session, err := openIDAuthService.NewStore().Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
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

// OnboardCustomer enables a logged in user to update their profile details
// @Summary Get profile enables a logged in user to update their profile details
// @Description Get profile enables a logged in user to update their profile details
// @Tags  Customers
// @Accept  json
// @Produce  json
// @Param customer body models.Onboarding true "Onboard User"
// @Success 200 {object} models.Onboarding{}
// @Router /auth/onboard [patch]
func (*customercontroller) OnboardCustomer(w http.ResponseWriter, r *http.Request) {
	var customerbody models.Customer

	err := json.NewDecoder(r.Body).Decode(&customerbody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	err = openIDAuthService.InitSession()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	session, err := openIDAuthService.NewStore().Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	profile := session.Values["profile"]
	email := fmt.Sprintf("%v", profile.(map[string]interface{})["email"])

	customerbody.Status = "active"
	customerbody.Code = "123"

	fmt.Printf("%+v\n", email)

	err = customerService.ValidateCustomerOnboarding(&customerbody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", err.Error())
		return
	}

	// TODO: Fix this to include email in header params
	// if customerbody.Email != email {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	utils.ResponseHelper(w, "401", "Only owner allowed to edit profile")
	// 	return
	// }

	customer, err := customerService.OnboardCustomer(email, customerbody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	utils.ResponseWithDataHelper(w, "200", "customer details updated successfully", customer)
}
