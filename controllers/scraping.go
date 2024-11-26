package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AD0791/SO/scraper/services"
	"github.com/julienschmidt/httprouter"
)

type ScrapingController struct {
	ScraperService *services.ScraperService
}

func NewScrapingController(scraperService *services.ScraperService) *ScrapingController {
	return &ScrapingController{
		ScraperService: scraperService,
	}
}

// @Summary Scrape StackOverflow Questions
// @Description Retrieves the latest StackOverflow questions and returns them as a JSON response
// @Tags Scraping
// @Success 200 {object} models.QuestionsResponse
// @Failure 500 {object} string
// @Router /scrape/stackoverflow [get]
func (c *ScrapingController) ScrapeStackOverflowQuestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	questionsResponse, err := c.ScraperService.ScrapeStackOverflowQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questionsResponse)
}
