package utils

import (
	"encoding/json"
	"net/http"
)

//MessageResponse struct declaration
type MessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//MessageWithData struct declaration
type MessageWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    dt     `json:"data"`
}

type dt interface {
}

// ResponseHelper Declaration
func ResponseHelper(w http.ResponseWriter, status, message string) {
	result := MessageResponse{}
	result.Status = status
	result.Message = message
	json.NewEncoder(w).Encode(result)
}

// ResponseWithDataHelper Declaration
func ResponseWithDataHelper(w http.ResponseWriter, status, message string, data dt) {
	result := MessageWithData{}
	result.Status = status
	result.Message = message
	result.Data = data
	json.NewEncoder(w).Encode(result)
}
