package config

import "os"

type DbConfig struct {
	User     string
	Password string
	DbName   string
	Host     string
	Port     string
}

type Config struct {
	DB DbConfig
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		DB: DbConfig{
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PWD", ""),
			DbName:   getEnv("DB_NAME", ""),
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
