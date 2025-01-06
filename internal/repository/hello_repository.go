package repository

import (
	"gin-samples/internal/domain"
	"gorm.io/gorm"
)

// HelloRepository extends CrudRepository with additional methods
type HelloRepository interface {
	CrudRepository[domain.Greeting, uint]
	ExistsByMessage(message string) (bool, error)
}

type helloRepositoryImpl struct {
	*BaseRepository[domain.Greeting, uint]
}

// NewHelloRepository creates a new instance of HelloRepository
func NewHelloRepository(db *gorm.DB) HelloRepository {
	return &helloRepositoryImpl{
		BaseRepository: NewBaseRepository[domain.Greeting, uint](db),
	}
}

func (r *helloRepositoryImpl) ExistsByMessage(message string) (bool, error) {
	var count int64
	if err := r.db.Model(&domain.Greeting{}).Where("message = ?", message).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
