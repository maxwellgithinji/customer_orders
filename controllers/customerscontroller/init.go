package customerscontroller

import (
	"net/http"

	service "github.com/maxwellgithinji/customer_orders/services"
)

type CustomerController interface {
	GetCustomers(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
}

type customercontroller struct{}

var (
	customerService   service.CustomerService
	openIDAuthService service.OpenIdAuthService
)

func NewCustomerController(service service.CustomerService, openIdAuth service.OpenIdAuthService) CustomerController {
	customerService = service
	openIDAuthService = openIdAuth
	return &customercontroller{}
}
