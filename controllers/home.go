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

func (*HomeController) GetHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := models.ApiResponse{
		Message: "Welcome to GoScraper API",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
