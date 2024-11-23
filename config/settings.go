package config

type server struct {
	Port       int    `mapstructure:"port"`
	ProjectEnv string `mapstructure:"project_env"`
	RootPath   string `mapstructure:"root_path"` // Fixed typo in tag
}

type logs struct {
	LogLevel string `mapstructure:"log_level"`
}

type swagDoc struct {
	EnableSwagger bool `mapstructure:"enable"`
}

type Config struct {
	AppName     string  `mapstructure:"app_name"`
	Server      server  `mapstructure:"server"`
	SwagDoc     swagDoc `mapstructure:"swaggerDoc"` // Fixed typo in tag
	Logs        logs    `mapstructure:"logs"`
	DatabaseURL string  `mapstructure:"database_url"` // Changed to lowercase for consistency
}
