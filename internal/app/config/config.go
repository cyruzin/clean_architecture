package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config is a configuration struct that contains
// all environment variables of the app.
type Config struct {
	EnvMode           string        `envconfig:"ENV_MODE" default:"development" required:"true"`
	Port              string        `envconfig:"PORT" default:"8000" required:"true"`
	ReadTimeOut       time.Duration `envconfig:"READ_TIMEOUT" default:"5s" required:"true"`
	ReadHeaderTimeOut time.Duration `envconfig:"READ_HEADER_TIMEOUT" default:"5s" required:"true"`
	WriteTimeOut      time.Duration `envconfig:"WRITE_TIMEOUT" default:"5s" required:"true"`
	IdleTimeOut       time.Duration `envconfig:"IDLE_TIMEOUT" default:"60s" required:"true"`
}

// Load loads the app the configuration based
// in the environment variables.
func Load() *Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return &cfg
}
