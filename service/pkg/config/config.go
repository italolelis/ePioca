package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Specification for basic configurations
type Specification struct {
	Port     int `envconfig:"PORT"`
	LogLevel int `envconfig:"LOG_LEVEL"`
	Database Database
}

// Database holds the configuration for a database
type Database struct {
	DSN            string `envconfig:"DATABASE_DSN"`
	MigrationsPath string `envconfig:"MIGRATIONS_PATH"`
}

//Load loads configuration from environment variables
func Load() (*Specification, error) {
	var config Specification

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
