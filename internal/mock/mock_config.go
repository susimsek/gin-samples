package mock

import (
	"gin-samples/config"
	"time"
)

// MockConfig returns a mock configuration for testing.
func MockConfig() *config.Config {
	return &config.Config{
		ServerPort:    "8080",
		TokenDuration: time.Minute * 30,
	}
}
