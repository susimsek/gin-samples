package repository

import (
	"gin-samples/internal/domain"
	"gorm.io/gorm"
)

type HelloRepository interface {
	SaveGreeting(greeting domain.Greeting) (domain.Greeting, error)
	GetAllGreetings() ([]domain.Greeting, error)
	FindByMessage(message string) (domain.Greeting, error)
	ExistsByMessage(message string) (bool, error)
}

type helloRepositoryImpl struct {
	db *gorm.DB
}

func NewHelloRepository(db *gorm.DB) HelloRepository {
	return &helloRepositoryImpl{db: db}
}

func (r *helloRepositoryImpl) SaveGreeting(greeting domain.Greeting) (domain.Greeting, error) {
	if err := r.db.Create(&greeting).Error; err != nil {
		return domain.Greeting{}, err
	}
	return greeting, nil
}

func (r *helloRepositoryImpl) GetAllGreetings() ([]domain.Greeting, error) {
	var greetings []domain.Greeting
	if err := r.db.Find(&greetings).Error; err != nil {
		return nil, err
	}
	return greetings, nil
}

func (r *helloRepositoryImpl) FindByMessage(message string) (domain.Greeting, error) {
	var greeting domain.Greeting
	if err := r.db.Where("message = ?", message).First(&greeting).Error; err != nil {
		return domain.Greeting{}, err
	}
	return greeting, nil
}

func (r *helloRepositoryImpl) ExistsByMessage(message string) (bool, error) {
	var count int64
	if err := r.db.Model(&domain.Greeting{}).Where("message = ?", message).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
