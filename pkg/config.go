// wf-dba: pkg/config.go
package pkg

import (
	"log"
	"os"
	"strconv"
)

// Config holds the configuration values for the application
type Config struct {
	Port     string
	LogLevel string
}

// LoadConfig loads configuration values from environment variables
func LoadConfig() *Config {
	return &Config{
		Port:     getEnv("APP_PORT", "8083"),
		LogLevel: getEnv("LOG_LEVEL", "INFO"),
	}
}

// Helper functions to read environment variables with fallbacks
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(name); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
		log.Printf("Invalid integer for %s: %v", name, valueStr)
	}
	return defaultValue
}
