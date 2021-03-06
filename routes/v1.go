package routes

import (
	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/customer_orders/controllers/authcontroller"
	"github.com/maxwellgithinji/customer_orders/controllers/customerscontroller"
	"github.com/maxwellgithinji/customer_orders/controllers/itemscontroller"
	"github.com/maxwellgithinji/customer_orders/controllers/orderscontroller"
	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/middlewares"
	"github.com/maxwellgithinji/customer_orders/services/customerservice"
	"github.com/maxwellgithinji/customer_orders/services/itemservice"
	"github.com/maxwellgithinji/customer_orders/services/openidauthservice"
	"github.com/maxwellgithinji/customer_orders/services/orderservice"
	"github.com/maxwellgithinji/customer_orders/utils/sms"
)

var (
	// SMS
	message sms.SMS = sms.NewSMS()

	// Database
	database databases.Database = databases.NewDatabase()

	// Middlewares
	authmiddleware middlewares.AuthMiddleware = middlewares.NewAuthMiddleware(openIDAuthService)

	// Auth
	openIDAuthService    openidauthservice.OpenIdAuthService = openidauthservice.NewOpenIdAuthService()
	openIDAuthController authcontroller.OpenIDAuthController = authcontroller.NewOpenIdAuthController(openIDAuthService, customerService)

	// Customer
	customerTable      databases.CustomerTable                = databases.NewCustomersTable(database)
	customerService    customerservice.CustomerService        = customerservice.NewCustomerService(customerTable)
	customerController customerscontroller.CustomerController = customerscontroller.NewCustomerController(customerService, openIDAuthService)

	// Items
	itemTable      databases.ItemTable            = databases.NewItemsTable(database)
	Itemservice    itemservice.ItemService        = itemservice.NewItemService(itemTable)
	itemcontroller itemscontroller.ItemController = itemscontroller.NewItemController(Itemservice)

	// Orders
	OrderTable      databases.OrderTable             = databases.NewOrdersTable(databases.DB)
	OrderService    orderservice.OrderService        = orderservice.NewOrderService(OrderTable)
	OrderController orderscontroller.OrderController = orderscontroller.NewOrderController(OrderService, customerService, Itemservice, openIDAuthService, message)
)

func apiV1(api *mux.Router) {
	var api1 = api.PathPrefix("/v1").Subrouter()
	/*
		Open Routes
	*/

	// Index
	api1.HandleFunc("/", index).Methods("GET")

	// Items
	api1.HandleFunc("/items", itemcontroller.GetItems).Methods("GET")

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
	a.HandleFunc("/onboard", customerController.OnboardCustomer).Methods("PATCH")

	// Items
	a.HandleFunc("/item", itemcontroller.CreateItem).Methods("POST")
	a.HandleFunc("/delete/item/{id}", itemcontroller.DeleteItem).Methods("DELETE")

	// Orders
	a.HandleFunc("/orders", OrderController.GetOrders).Methods("GET")
	a.HandleFunc("/orders", OrderController.CreateOrder).Methods("POST")
	a.HandleFunc("/currentuser/orders", OrderController.FindCurrentUserOrders).Methods("GET")
}
