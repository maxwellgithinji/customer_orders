package orderscontroller

import (
	"net/http"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/services/customerservice"
	"github.com/maxwellgithinji/customer_orders/services/itemservice"
	"github.com/maxwellgithinji/customer_orders/services/openidauthservice"
	"github.com/maxwellgithinji/customer_orders/services/orderservice"
)

type OrderController interface {
	GetOrders(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
	FindCurrentUserOrders(w http.ResponseWriter, r *http.Request)
}

type ordercontroller struct{}

var (
	// DB init
	OrderTable    databases.OrderTable    = databases.NewOrdersTable(databases.DB)
	CustomerTable databases.CustomerTable = databases.NewCustomersTable(databases.DB)
	ItemTable     databases.ItemTable     = databases.NewItemsTable(databases.DB)

	// Services init
	OrderService      orderservice.OrderService           = orderservice.NewOrderService(OrderTable)
	CustomerService   customerservice.CustomerService     = customerservice.NewCustomerService(CustomerTable)
	ItemService       itemservice.ItemService             = itemservice.NewItemService(ItemTable)
	OpenIDAuthService openidauthservice.OpenIdAuthService = openidauthservice.NewOpenIdAuthService()
)

func NewOrderController(
	order orderservice.OrderService,
	customer customerservice.CustomerService,
	item itemservice.ItemService,
	openid openidauthservice.OpenIdAuthService) OrderController {
	OrderService = order
	CustomerService = customer
	ItemService = item
	OpenIDAuthService = openid
	return &ordercontroller{}
}
