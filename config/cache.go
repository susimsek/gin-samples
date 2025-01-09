package config

import (
	"github.com/dgraph-io/ristretto"
	"log"
)

// InitCache initializes and returns a Ristretto cache instance.
func InitCache() *ristretto.Cache {
	// Create a new Ristretto cache with the specified configuration
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // Number of counters used to track access frequency
		MaxCost:     1 << 30, // Maximum memory usage (1 GB)
		BufferItems: 64,      // Performance optimization buffer
	})
	if err != nil {
		// Log an error and stop the application if cache initialization fails
		log.Fatalf("Failed to create cache: %v", err)
	}

	return cache
}
