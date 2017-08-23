package config

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
