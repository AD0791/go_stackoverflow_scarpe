package router

import (
	"net/http"

	"github.com/AD0791/SO/scraper/config"
	"github.com/AD0791/SO/scraper/controllers"
	"github.com/AD0791/SO/scraper/services"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	config             *config.Config
	router             *httprouter.Router
	homeController     *controllers.HomeController
	scrapingController *controllers.ScrapingController
}

func NewRouter(cfg *config.Config) *Router {
	scraperService := services.NewScraperService()
	scrapingController := controllers.NewScrapingController(scraperService)

	return &Router{
		config:             cfg,
		router:             httprouter.New(),
		homeController:     controllers.NewHomeController(),
		scrapingController: scrapingController,
	}
}

func (r *Router) Setup() http.Handler {
	// Register routes
	r.router.GET(r.config.Server.RootPath, r.homeController.GetHome)
	r.router.GET("/scrape/stackoverflow", r.scrapingController.ScrapeStackOverflowQuestions)

	// Swagger endpoint
	r.router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	return r.router
}
