package utils

import (
	"encoding/json"
	"net/http"

	"github.com/AD0791/SO/scraper/models"
)

// SendResponse sends a structured JSON response
func SendResponse(w http.ResponseWriter, statusCode int, message string, data interface{}, err string) {
	response := &models.ApiResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
