package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort string
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	_ = godotenv.Load(".env")

	envFile := ".env." + env
	_ = godotenv.Load(envFile)

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
