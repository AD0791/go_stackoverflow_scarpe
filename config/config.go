package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var AppConfig *Config

// TODO FIX THE PATH TO BE UNIX/UNIQUE
const (
	PROJECT_PATH  = "/Users/ad0791/Desktop/go_stackoverflow_scarpe"
	ENV_FILE_PATH = "/Users/ad0791/Desktop/go_stackoverflow_scarpe/.env"
)

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.AddConfigPath(PROJECT_PATH) // look for the config file in the current directory

	if err := godotenv.Load(ENV_FILE_PATH); err != nil {
		log.Println("Warning: Could not load .env file, falling back to environment variables")
		log.Println("Warning: Could not load .env file, falling back to environment variables")
	}

	// Set default values
	v.SetDefault("app_name", "GoScraper")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.project_env", "development")
	v.SetDefault("server.root_path", "/")
	v.SetDefault("swaggerDoc.enable", true)
	v.SetDefault("logs.log_level", "debug")

	// Configure Viper for YAML
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	//v.SetEnvPrefix("APP")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Could not load configuration file: %v", err)
		return nil, fmt.Errorf("Could not load configuration file: %v", err)
	}

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	if err := v.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Could not build the AppConfig type: %v", err)
		return nil, fmt.Errorf("Could not build the AppConfig type: %v", err)

	}

	// Handle DATABASE_URL separately since it's a direct environment variable
	if dbURL := v.GetString("DATABASE_URL"); dbURL != "" {
		AppConfig.DatabaseURL = dbURL
	}

	log.Printf("Configuration loaded: %+v", AppConfig)
	return AppConfig, nil
}
