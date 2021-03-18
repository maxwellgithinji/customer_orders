package routes

import (
	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/customer_orders/controllers/authcontroller"
	"github.com/maxwellgithinji/customer_orders/controllers/customerscontroller"
	"github.com/maxwellgithinji/customer_orders/controllers/itemscontroller"
	"github.com/maxwellgithinji/customer_orders/middlewares"
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
	api1.HandleFunc("/login", authcontroller.Login).Methods("GET")
	api1.HandleFunc("/logout", authcontroller.Logout).Methods("POST")
	api1.HandleFunc("/callback", authcontroller.Callback).Methods("GET")

	/*
		Authenticated routes
	*/
	a := api1.PathPrefix("/auth").Subrouter()
	a.Use(middlewares.IsAuthenticated)

	// Customers
	a.HandleFunc("/profile", customerscontroller.Profile).Methods("GET")
	a.HandleFunc("/customers", customerscontroller.GetCustomers).Methods("GET")

}
