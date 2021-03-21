package authcontroller

import (
	"net/http"

	service "github.com/maxwellgithinji/customer_orders/services"
)

type OpenIDAuthController interface {
	Callback(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}
type authcontroller struct{}

var (
	openIDAuthService service.OpenIdAuthService
)

func NewOpenIdAuthController(service service.OpenIdAuthService) OpenIDAuthController {
	openIDAuthService = service
	return &authcontroller{}
}
