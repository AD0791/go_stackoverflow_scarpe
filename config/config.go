package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

//TODO fix viper -> viper must see root_path where config files are.
//TODO fix viper -> .env
//TODO fix viper -> yml
//TODO fix viper -> fix the main

var AppConfig *Config

func LoadConfig() (*Config, error) {
	RootPath := viper.GetString("root_path")
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file, falling back to environment variables")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(RootPath) // look for the config file in the current directory

	// Read environment variables as a fallback
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Could not load configuration file: %v", err)
	}

	// Map environment variables (like DATABASE_URL)
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("Warning: DATABASE_URL is not set in .env or environment")
	}

	AppConfig := &Config{
		AppName:       viper.GetString("app_name"),
		Port:          viper.GetInt("port"),
		Environment:   viper.GetString("env"),
		EnableSwagger: viper.GetBool("enable_swagger"),
		LogLevel:      viper.GetString("log_level"),
		RootPath:      viper.GetString("root_path"),
		DatabaseURL:   dbURL, // Read DATABASE_URL from .env
	}

	log.Printf("Configuration loaded: %+v", AppConfig)
	return AppConfig, nil
}
