package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AD0791/SO/scraper/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("failed to load configuration: %v", err)
	}

	// Use configuration values
	log.Printf("Starting %s in %s mode on port %d",
		cfg.AppName,
		cfg.Server.ProjectEnv,
		cfg.Server.Port,
	)

	// Example of accessing Database URL
	log.Printf("Database URL: %s", cfg.DatabaseURL)

	http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), nil)
}
