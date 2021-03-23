package authcontroller

import (
	"net/http"

	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/services/customerservice"
	"github.com/maxwellgithinji/customer_orders/services/openidauthservice"
)

type OpenIDAuthController interface {
	Callback(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}
type authcontroller struct{}

var (
	openIDAuthService openidauthservice.OpenIdAuthService = openidauthservice.NewOpenIdAuthService()
	customerTable     databases.CustomerTable             = databases.NewCustomersTable(databases.DB)
	customerService   customerservice.CustomerService     = customerservice.NewCustomerService(customerTable)
)

func NewOpenIdAuthController(service openidauthservice.OpenIdAuthService, customer databases.CustomerTable) OpenIDAuthController {
	openIDAuthService = service
	customerTable = customer
	return &authcontroller{}
}
