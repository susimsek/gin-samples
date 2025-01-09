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

// UserRepository defines additional methods for User-specific queries
type UserRepository interface {
	CrudRepository[domain.User, string]
	FindByUsername(username string) (util.Optional[domain.User], error)
	FindByEmail(email string) (util.Optional[domain.User], error)
}

type userRepositoryImpl struct {
	*BaseRepository[domain.User, string]
	cacheManager *cache.CacheManager
	db           *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB, cacheManager *cache.CacheManager) UserRepository {
	return &userRepositoryImpl{
		BaseRepository: NewBaseRepository[domain.User, string](db, cacheManager, "user"),
		cacheManager:   cacheManager,
		db:             db,
	}
}

// FindByUsername retrieves a user by their username, including roles and caches the result
func (r *userRepositoryImpl) FindByUsername(username string) (util.Optional[domain.User], error) {
	cacheKey := fmt.Sprintf("userByUsername:%s", username)

	// Check the cache first
	if cachedValue, found := r.cacheManager.Get(cacheKey); found {
		return util.Optional[domain.User]{Value: cachedValue.(*domain.User)}, nil
	}

	// If not in cache, query the database
	var user domain.User
	err := r.db.Preload("Roles.Role").Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util.Optional[domain.User]{Value: nil}, nil
		}
		return util.Optional[domain.User]{}, err
	}

	// Cache the result with a 1-hour TTL
	r.cacheManager.Set(cacheKey, &user, 1*time.Hour)

	return util.Optional[domain.User]{Value: &user}, nil
}

// FindByEmail retrieves a user by their email without roles and caches the result
func (r *userRepositoryImpl) FindByEmail(email string) (util.Optional[domain.User], error) {
	cacheKey := fmt.Sprintf("userByEmail:%s", email)

	// Check the cache first
	if cachedValue, found := r.cacheManager.Get(cacheKey); found {
		return util.Optional[domain.User]{Value: cachedValue.(*domain.User)}, nil
	}

	// If not in cache, query the database
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util.Optional[domain.User]{Value: nil}, nil
		}
		return util.Optional[domain.User]{}, err
	}

	// Cache the result with a 1-hour TTL
	r.cacheManager.Set(cacheKey, &user, 1*time.Hour)

	return util.Optional[domain.User]{Value: &user}, nil
}
