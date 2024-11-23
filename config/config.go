package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var AppConfig *Config

func getProjectRoot() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// If we're in a subdirectory (like /config), we need to go up one level
	if filepath.Base(pwd) == "config" {
		return filepath.Dir(pwd)
	}

	return pwd
}

func LoadConfig() (*Config, error) {
	projectRoot := getProjectRoot()

	v := viper.New()

	// Add project root as config path
	v.AddConfigPath(projectRoot)
	envPath := filepath.Join(projectRoot, ".env")
	if err := godotenv.Load(envPath); err != nil {
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
