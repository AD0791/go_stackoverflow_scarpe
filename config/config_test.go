package config_test

import (
	"os"
	"testing"

	"github.com/AD0791/SO/scraper/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set up temporary .env file
	os.Setenv("DATABASE_URL", "test_database_url")

	// Load configuration
	cfg, err := config.LoadConfig()
	assert.NoError(t, err, "Loading configuration should not return an error")
	assert.Equal(t, "GoScraper", cfg.AppName, "AppName should match config.yml")
	assert.Equal(t, 8080, cfg.Port, "Port should match config.yml")
	assert.Equal(t, "development", cfg.Environment, "Environment should match config.yml")
	assert.Equal(t, "test_database_url", cfg.DatabaseURL, "DATABASE_URL should match the .env variable")
	assert.Equal(t, "./", cfg.RootPath, "RootPath should match config.yml")
	assert.True(t, cfg.EnableSwagger, "EnableSwagger should be true as per config.yml")
}
