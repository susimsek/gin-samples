package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
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
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
