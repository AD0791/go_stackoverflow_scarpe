package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var AppConfig *Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.AddConfigPath("/Users/ad0791/Desktop/go_stackoverflow_scarpe") // look for the config file in the current directory

	if err := godotenv.Load("/Users/ad0791/Desktop/go_stackoverflow_scarpe/.env"); err != nil {
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
