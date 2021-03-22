package routes

import (
	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/customer_orders/controllers/authcontroller"
	"github.com/maxwellgithinji/customer_orders/controllers/customerscontroller"
	"github.com/maxwellgithinji/customer_orders/controllers/itemscontroller"
	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/middlewares"
	service "github.com/maxwellgithinji/customer_orders/services"
)

var (

	// Auth
	openIDAuthService    service.OpenIdAuthService           = service.NewOpenIdAuthService()
	openIDAuthController authcontroller.OpenIDAuthController = authcontroller.NewOpenIdAuthController(openIDAuthService)

	// Middlewares
	authmiddleware middlewares.AuthMiddleware = middlewares.NewAuthMiddleware(openIDAuthService)
	// Customer
	customerTable      databases.CustomerTable                = databases.NewCustomersTable()
	customerService    service.CustomerService                = service.NewCustomerService(customerTable)
	customerController customerscontroller.CustomerController = customerscontroller.NewCustomerController(customerService, openIDAuthService)
)

func apiV1(api *mux.Router) {
	var api1 = api.PathPrefix("/v1").Subrouter()
	/*
		Open Routes
	*/

	// Index
	api1.HandleFunc("/", index).Methods("GET")

	// Items
	api1.HandleFunc("/items", itemscontroller.GetItems).Methods("GET")

	// Authorization/Authentication
	api1.HandleFunc("/callback", openIDAuthController.Callback).Methods("Get")
	api1.HandleFunc("/login", openIDAuthController.Login).Methods("Get")
	api1.HandleFunc("/logout", openIDAuthController.Logout).Methods("POST")

	/*
		Authenticated routes
	*/
	a := api1.PathPrefix("/auth").Subrouter()
	a.Use(authmiddleware.IsAuthenticated)

	// Customers
	a.HandleFunc("/profile", customerController.Profile).Methods("GET")
	a.HandleFunc("/customers", customerController.GetCustomers).Methods("GET")

}
