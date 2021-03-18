package routes

import (
	"encoding/json"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/utils"
)

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	msg := utils.MessageResponse{
		Message: "Welcome to customer orders, login to continue",
	}
	json.NewEncoder(w).Encode(msg)
}

func customer(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	msg := utils.MessageResponse{
		Message: "Welcome customer",
	}
	json.NewEncoder(w).Encode(msg)
}
