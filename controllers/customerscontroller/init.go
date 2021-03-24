package customerscontroller

import (
	"net/http"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/services/customerservice"
	"github.com/maxwellgithinji/customer_orders/services/openidauthservice"
)

type CustomerController interface {
	GetCustomers(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	OnboardCustomer(w http.ResponseWriter, r *http.Request)
}

type customercontroller struct{}

var (
	CustomerTable     databases.CustomerTable             = databases.NewCustomersTable(databases.DB)
	customerService   customerservice.CustomerService     = customerservice.NewCustomerService(CustomerTable)
	openIDAuthService openidauthservice.OpenIdAuthService = openidauthservice.NewOpenIdAuthService()
)

func NewCustomerController(service customerservice.CustomerService, openIdAuth openidauthservice.OpenIdAuthService) CustomerController {
	customerService = service
	openIDAuthService = openIdAuth
	return &customercontroller{}
}
