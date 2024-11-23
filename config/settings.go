package config

type Config struct {
	AppName       string
	Port          int
	Environment   string
	EnableSwagger bool
	LogLevel      string
	RootPath      string
	DatabaseURL   string
}
