package router

import (
	"net/http"

	"github.com/AD0791/SO/scraper/config"
	"github.com/AD0791/SO/scraper/controllers"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	config   *config.Config
	router   *httprouter.Router
	homeCtrl *controllers.HomeController
}

func NewRouter(cfg *config.Config) *Router {
	return &Router{
		config:   cfg,
		router:   httprouter.New(),
		homeCtrl: controllers.NewHomeController(),
	}
}

func (r *Router) Setup() *httprouter.Router {
	// Register routes
	r.router.GET(r.config.Server.RootPath, r.homeCtrl.GetHome)

	// Swagger endpoint
	r.router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	return r.router
}
