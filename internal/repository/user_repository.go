package repository

import (
	"errors"
	"gin-samples/internal/domain"
	"gin-samples/internal/util"
	"gorm.io/gorm"
)

// UserRepository defines additional methods for User-specific queries
type UserRepository interface {
	CrudRepository[domain.User, string]
	FindByUsername(username string) (util.Optional[domain.User], error)
	FindByEmail(email string) (util.Optional[domain.User], error)
}

type userRepositoryImpl struct {
	*BaseRepository[domain.User, string]
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		BaseRepository: NewBaseRepository[domain.User, string](db),
		db:             db,
	}
}

// FindByUsername retrieves a user by their username, including roles
func (r *userRepositoryImpl) FindByUsername(username string) (util.Optional[domain.User], error) {
	var user domain.User
	err := r.db.Preload("Roles.Role").Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util.Optional[domain.User]{Value: nil}, nil
		}
		return util.Optional[domain.User]{}, err
	}
	return util.Optional[domain.User]{Value: &user}, nil
}

// FindByEmail retrieves a user by their email without roles
func (r *userRepositoryImpl) FindByEmail(email string) (util.Optional[domain.User], error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util.Optional[domain.User]{Value: nil}, nil
		}
		return util.Optional[domain.User]{}, err
	}
	return util.Optional[domain.User]{Value: &user}, nil
}
