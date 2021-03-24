package authcontroller

import (
	"github.com/maxwellgithinji/customer_orders/services/customerservice"
	"github.com/maxwellgithinji/customer_orders/services/openidauthservice"
)

var (
	customerServiceTest      customerservice.CustomerService
	openIDAuthServiceTest    openidauthservice.OpenIdAuthService = openidauthservice.NewOpenIdAuthService()
	openIDAuthControllerTest OpenIDAuthController                = NewOpenIdAuthController(openIDAuthServiceTest, customerServiceTest)
)
