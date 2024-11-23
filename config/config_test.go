package config_test

import (
	"os"
	"testing"

	"github.com/AD0791/SO/scraper/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set up environment variables for testing
	os.Setenv("DATABASE_URL", "test_database_url")

	// Load configuration
	cfg, err := config.LoadConfig()

	// Assert no error occurred
	assert.NoError(t, err, "Loading configuration should not return an error")

	// Assert configuration values
	assert.Equal(t, "GoScraper", cfg.AppName, "AppName should match config.yml")
	assert.NotNil(t, cfg.Server, "Server config should not be nil")
	assert.Equal(t, 8080, cfg.Server.Port, "Port should match config.yml")
	assert.Equal(t, "development", cfg.Server.ProjectEnv, "Environment should match config.yml")
	assert.Equal(t, "/", cfg.Server.RootPath, "RootPath should match config.yml")
	assert.NotNil(t, cfg.SwagDoc, "SwagDoc config should not be nil")
	assert.True(t, cfg.SwagDoc.EnableSwagger, "EnableSwagger should be true as per config.yml")
	assert.Equal(t, "test_database_url", cfg.DatabaseURL, "DATABASE_URL should match the environment variable")

	// Clean up
	os.Unsetenv("DATABASE_URL")
}
