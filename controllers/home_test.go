package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AD0791/SO/scraper/models"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetHome(t *testing.T) {
	// Create a new router instance
	router := httprouter.New()

	// Create a new controller instance
	homeCtrl := NewHomeController()

	// Register the route
	router.GET("/", homeCtrl.GetHome)

	// Create a test request
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse response
	var response models.ApiResponse
	err := json.NewDecoder(rr.Body).Decode(&response)

	// Assert response
	assert.NoError(t, err)
	assert.Equal(t, "Welcome to GoScraper API", response.Message)
}
