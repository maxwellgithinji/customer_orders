package authcontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogout(t *testing.T) {
	req, _ := http.NewRequest("GET", "/login", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(openIDAuthControllerTest.Logout)
	handler.ServeHTTP(rr, req)

	// status := rr.Result().StatusCode

	// if status != 200 {
	// 	t.Errorf("handler returned a wrong status, got: %v want %v", status, 200)
	// }
}
