package authcontroller

import (
	"github.com/maxwellgithinji/customer_orders/databases"
	"github.com/maxwellgithinji/customer_orders/services/openidauthservice"
)

var (
	customerTableTest        databases.CustomerTable
	openIDAuthServiceTest    openidauthservice.OpenIdAuthService = openidauthservice.NewOpenIdAuthService()
	openIDAuthControllerTest OpenIDAuthController                = NewOpenIdAuthController(openIDAuthServiceTest, customerTableTest)
)
