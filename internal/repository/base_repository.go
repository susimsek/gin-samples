package repository

import (
	"errors"
	"fmt"
	"gin-samples/internal/util"
	"gorm.io/gorm"
)

// CrudRepository defines the standard CRUD operations
type CrudRepository[T any, ID any] interface {
	Save(entity T) (T, error)
	FindAll() ([]T, error)
	FindByID(id ID) (util.Optional[T], error) // Optional return type
	DeleteByID(id ID) error
}

// BaseRepository provides generic CRUD operations for any entity type T with ID type
type BaseRepository[T any, ID any] struct {
	db *gorm.DB
}

// NewBaseRepository creates a new instance of BaseRepository
func NewBaseRepository[T any, ID any](db *gorm.DB) *BaseRepository[T, ID] {
	return &BaseRepository[T, ID]{db: db}
}

// Save creates or updates an entity
func (r *BaseRepository[T, ID]) Save(entity T) (T, error) {
	if err := r.db.Save(&entity).Error; err != nil {
		return *new(T), fmt.Errorf("failed to save entity: %w", err)
	}
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

// FindByID retrieves an entity by its ID and returns Optional
func (r *BaseRepository[T, ID]) FindByID(id ID) (util.Optional[T], error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return an empty Optional if record is not found
			return util.EmptyOptional[T](), nil
		}
		return util.Optional[T]{}, fmt.Errorf("failed to fetch entity by ID: %w", err)
	}
	return util.Optional[T]{Value: &entity}, nil
}

// DeleteByID deletes an entity by its ID
func (r *BaseRepository[T, ID]) DeleteByID(id ID) error {
	if err := r.db.Delete(new(T), id).Error; err != nil {
		return fmt.Errorf("failed to delete entity by ID: %w", err)
	}
	return nil
}
