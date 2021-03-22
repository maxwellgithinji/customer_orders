package authcontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	service "github.com/maxwellgithinji/customer_orders/services"
)

var (
	openIDAuthServiceLoginTest    service.OpenIdAuthService = service.NewOpenIdAuthService()
	openIDAuthControllerLoginTest OpenIDAuthController      = NewOpenIdAuthController(openIDAuthServiceTest)
)

func TestLogin(t *testing.T) {
	req, _ := http.NewRequest("GET", "/login", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(openIDAuthControllerLoginTest.Login)
	handler.ServeHTTP(rr, req)

	status := rr.Result().StatusCode

	if status != 200 {
		t.Errorf("handler returned a wrong status, got: %v want %v", status, 200)
	}
}
