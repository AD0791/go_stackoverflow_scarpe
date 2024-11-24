package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AD0791/SO/scraper/config"
	_ "github.com/AD0791/SO/scraper/docs" // Import generated docs
	"github.com/AD0791/SO/scraper/router"
)

// @title GoScraper API
// @version 1.0
// @description API documentation for GoScraper.
// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Initialize router
	r := router.NewRouter(cfg)

	// Setup routes
	handler := r.Setup()

	// Create server address
	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	// Log startup information
	log.Printf("Starting %s in %s mode on port %d",
		cfg.AppName,
		cfg.Server.ProjectEnv,
		cfg.Server.Port,
	)
	log.Printf("API root path: %s", cfg.Server.RootPath)

	// Start the server
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
