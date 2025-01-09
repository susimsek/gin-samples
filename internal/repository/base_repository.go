package repository

import (
	"errors"
	"fmt"
	"gin-samples/internal/cache"
	"gin-samples/internal/domain"
	"gin-samples/internal/util"
	"gorm.io/gorm"
	"time"
)

// CrudRepository defines the standard CRUD operations
type CrudRepository[T domain.Identifiable, ID any] interface {
	Save(entity T) (T, error)
	FindAll() ([]T, error)
	FindByID(id ID) (util.Optional[T], error)
	DeleteByID(id ID) error
}

// BaseRepository provides generic CRUD operations for any entity type T with ID type
type BaseRepository[T domain.Identifiable, ID any] struct {
	db           *gorm.DB
	cacheManager *cache.CacheManager
	cacheName    string
}

// NewBaseRepository creates a new instance of BaseRepository
func NewBaseRepository[T domain.Identifiable, ID any](db *gorm.DB, cacheManager *cache.CacheManager, cacheName string) *BaseRepository[T, ID] {
	return &BaseRepository[T, ID]{
		db:           db,
		cacheManager: cacheManager,
		cacheName:    cacheName,
	}
}

// Save creates or updates an entity and updates the cache
func (r *BaseRepository[T, ID]) Save(entity T) (T, error) {
	if err := r.db.Save(&entity).Error; err != nil {
		return *new(T), fmt.Errorf("failed to save entity: %w", err)
	}

	// Cache the entity with a 1-hour TTL using the entity's ID
	cacheKey := fmt.Sprintf("%s:%v", r.cacheName, entity.GetID())
	r.cacheManager.Set(cacheKey, &entity, 1*time.Hour)

	return entity, nil
}

// FindAll retrieves all entities
func (r *BaseRepository[T, ID]) FindAll() ([]T, error) {
	var entities []T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch entities: %w", err)
	}
	return entities, nil
}

// FindByID retrieves an entity by its ID and caches the result
func (r *BaseRepository[T, ID]) FindByID(id ID) (util.Optional[T], error) {
	// Build the cache key using the entity's ID
	var entity T
	cacheKey := fmt.Sprintf("%s:%v", r.cacheName, id)

	// Check the cache first
	if cachedValue, found := r.cacheManager.Get(cacheKey); found {
		return util.Optional[T]{Value: cachedValue.(*T)}, nil
	}

	// If not in cache, query the database
	if err := r.db.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util.EmptyOptional[T](), nil
		}
		return util.Optional[T]{}, fmt.Errorf("failed to fetch entity by ID: %w", err)
	}

	// Cache the result with a 1-hour TTL
	r.cacheManager.Set(cacheKey, &entity, 1*time.Hour)

	return util.Optional[T]{Value: &entity}, nil
}

// DeleteByID deletes an entity by its ID and removes it from the cache
func (r *BaseRepository[T, ID]) DeleteByID(id ID) error {
	// Build the cache key using the entity's ID
	cacheKey := fmt.Sprintf("%s:%v", r.cacheName, id)

	if err := r.db.Delete(new(T), id).Error; err != nil {
		return fmt.Errorf("failed to delete entity by ID: %w", err)
	}

	// Remove from the cache
	r.cacheManager.Delete(cacheKey)

	return nil
}
