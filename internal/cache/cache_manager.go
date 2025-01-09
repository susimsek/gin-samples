package cache

import (
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
)

// CacheManager provides a centralized cache with TTL support
type CacheManager struct {
	cache *ristretto.Cache
	mutex sync.Mutex // Ensures thread-safe cache operations
}

// cacheItem wraps the value with an expiration time
type cacheItem struct {
	Value      interface{}
	Expiration time.Time
}

// NewCacheManager creates a new CacheManager
func NewCacheManager(cache *ristretto.Cache) *CacheManager {
	return &CacheManager{
		cache: cache,
	}
}

// Set adds a key-value pair to the cache with a specified TTL
func (cm *CacheManager) Set(key string, value interface{}, ttl time.Duration) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.cache.Set(key, cacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}, 1) // Adjust cost as needed
}

// Get retrieves a value from the cache and checks for TTL expiration
func (cm *CacheManager) Get(key string) (interface{}, bool) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cachedItem, found := cm.cache.Get(key)
	if !found {
		return nil, false
	}

	item := cachedItem.(cacheItem)
	if time.Now().After(item.Expiration) {
		// Remove expired item
		cm.cache.Del(key)
		return nil, false
	}

	return item.Value, true
}

// Delete removes a key-value pair from the cache
func (cm *CacheManager) Delete(key string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.cache.Del(key)
}

// Flush clears all items from the cache
func (cm *CacheManager) Flush() {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.cache.Clear()
}
