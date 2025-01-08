package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string
	TokenDuration time.Duration
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	commonEnvPath := filepath.Join("resources", "config", ".env")
	envFilePath := filepath.Join("resources", "config", ".env."+env)

	if err := godotenv.Overload(commonEnvPath, envFilePath); err != nil {
		log.Printf("Error loading .env files: %v", err)
	}

	return &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		TokenDuration: parseDuration("TOKEN_DURATION", "3600s"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// parseDuration parses a duration string from the environment or uses a default.
func parseDuration(key, defaultValue string) time.Duration {
	durationStr := getEnv(key, defaultValue)
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Printf("Invalid duration for %s: %s, using default: %s. Error: %v", key, durationStr, defaultValue, err)
		duration, _ = time.ParseDuration(defaultValue)
	}
	return duration
}
