// controllers/home.go
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

// GetHome handles the root route.
//
// @Summary Show Scraper Home Endpoint
// @Description Get the welcome message for the API
// @Tags Home
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ApiResponse
// @Router / [get]
func (*HomeController) GetHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := models.ApiResponse{
		Message: "Welcome to GoScraper API",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
