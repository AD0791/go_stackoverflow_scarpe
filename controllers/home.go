package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AD0791/SO/scraper/models"
	"github.com/julienschmidt/httprouter"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

// GetHome handles the root endpoint
// @Summary Get the home message
// @Description Returns a welcome message
// @Tags Home
// @Success 200 {object} models.ApiResponse
// @Failure 500 {object} models.ApiResponse
// @Router / [get]
func (*HomeController) GetHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := models.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "Welcome to GoScraper API",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
