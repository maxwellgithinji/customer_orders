package authcontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallback(t *testing.T) {
	req, _ := http.NewRequest("GET", "/callback", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(openIDAuthControllerTest.Callback)
	handler.ServeHTTP(rr, req)

	// status := rr.Result().StatusCode

	// if status != 200 {
	// 	t.Errorf("handler returned a wrong status, got: %v want %v", status, 200)
	// }
}
