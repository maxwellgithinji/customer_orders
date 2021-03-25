package customerscontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetCustomers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/auth/customers", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(customerControllerTest.GetCustomers)
	handler.ServeHTTP(rr, req)

	// status := rr.Result().StatusCode

	// if status != 200 {
	// 	t.Errorf("handler returned a wrong status, got: %v want %v", status, 200)
	// }
}
