package models

type ApiResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`  // Optional, for successful responses
	Error      string      `json:"error,omitempty"` // Optional, for error details
}
