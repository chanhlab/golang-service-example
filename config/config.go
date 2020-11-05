package config

import "github.com/ilyakaznacheev/cleanenv"

// Config is a application configuration structure
type Config struct {
	MySQL struct {
		Host              string `env:"MYSQL_HOST" env-description:"MySQL host"`
		DBName            string `env:"MYSQL_DATABASE" env-description:"MySQL Database name"`
		Username          string `env:"MYSQL_USERNAME" env-description:"MySQL user name"`
		Password          string `env:"MYSQL_PASSWORD" env-description:"MySQL user password"`
		MaxIDLEConnection int    `env:"MYSQL_MAX_IDLE_CONNECTION" env-description:"The maximum of number connections that should be kept all the time" env-default:"1000"`
		MaxOpenConnection int    `env:"MYSQL_MAX_OPEN_CONNECTION" env-description:"The maximum of open connections to the database" env-default:"100"`
	}

	Server struct {
		HTTPPort int `env:"HTTP_PORT" env-description:"The port number of REST API" env-default:"5001"`
		GRPCPort int `env:"GRPC_PORT" env-description:"The port number of gRPC API" env-default:"5000"`
	}

	Logger struct {
		LogLevel      int    `env:"LOG_LEVEL" env-description:"A Level is a logging priority" env-default:"0"`
		LogTimeFormat string `env:"LOG_TIME_FORMAT" env-description:"The log time format" env-default:"2006-01-02T15:04:05Z07:00"`
	}

	Redis struct {
		Host     string `env:"REDIS_HOST" env-description:"Redis hostname" env-default:"120.0.0.1"`
		Port     int    `env:"REDIS_PORT" env-description:"Redis port number" env-default:"6379"`
		Database int    `env:"REDIS_DATABASE" env-description:"Redis database" env-default:"0"`
	}
}

// AppConfig is application configurations
var AppConfig *Config

// NewConfig creates a new config
func NewConfig() {
	AppConfig = &Config{}
	err := cleanenv.ReadEnv(AppConfig)
	if err != nil {
		panic(err)
	}
}
