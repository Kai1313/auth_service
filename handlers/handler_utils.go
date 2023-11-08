// handler_utils.go
package handlers

import (
	"AuthService/models"
	"encoding/json"
	"net/http"
)

func handleError(w http.ResponseWriter, result bool, err error) {
	errorResponse := &models.ErrorResponse{
		Result: result,
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(errorResponse)
}

func handleResult(w http.ResponseWriter, result bool, data interface{}, message string) {
	response := &models.Response{
		Result: result,
		Message: message,
		Data: data,
	}
	json.NewEncoder(w).Encode(response)
}