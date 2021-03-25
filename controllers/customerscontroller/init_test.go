package customerscontroller

import (
	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/services/customerservice"
	"github.com/maxwellgithinji/customer_orders/services/openidauthservice"
)

var (
	CustomerTableTest      databases.CustomerTable             = databases.NewCustomersTable(databases.DB)
	customerServiceTest    customerservice.CustomerService     = customerservice.NewCustomerService(CustomerTable)
	openIDAuthServiceTest  openidauthservice.OpenIdAuthService = openidauthservice.NewOpenIdAuthService()
	customerControllerTest CustomerController                  = NewCustomerController(customerServiceTest, openIDAuthServiceTest)
)
