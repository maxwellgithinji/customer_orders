package itemscontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItemsUnauthenticatedDB(t *testing.T) {
	req, _ := http.NewRequest("GET", "/items", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(ItemControllerTest.GetItems)
	handler.ServeHTTP(rr, req)

	status := rr.Result().StatusCode

	if status != 500 {
		t.Errorf("handler returned a wrong status, got: %v want %v", status, 500)
	}
}
