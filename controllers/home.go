package controllers

import (
	"net/http"

	"github.com/AD0791/SO/scraper/utils"
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
	/* pattern can be replicate for a repositry */
	utils.SendResponse(w, http.StatusOK, "Welcome to GoScraper API", nil, "")
}
