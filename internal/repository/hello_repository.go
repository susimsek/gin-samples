package repository

import (
	"gin-samples/internal/entity"
	"gorm.io/gorm"
)

type HelloRepository interface {
	SaveGreeting(greeting *entity.Greeting) (*entity.Greeting, error)
	GetAllGreetings() ([]entity.Greeting, error)
	FindByMessage(message string) (*entity.Greeting, error)
	ExistsByMessage(message string) (bool, error)
}

type helloRepositoryImpl struct {
	db *gorm.DB
}

func NewHelloRepository(db *gorm.DB) HelloRepository {
	return &helloRepositoryImpl{db: db}
}

func (r *helloRepositoryImpl) SaveGreeting(greeting *entity.Greeting) (*entity.Greeting, error) {
	if err := r.db.Create(greeting).Error; err != nil {
		return nil, err
	}
	return greeting, nil
}

func (r *helloRepositoryImpl) GetAllGreetings() ([]entity.Greeting, error) {
	var greetings []entity.Greeting
	if err := r.db.Find(&greetings).Error; err != nil {
		return nil, err
	}
	return greetings, nil
}

func (r *helloRepositoryImpl) FindByMessage(message string) (*entity.Greeting, error) {
	var greeting entity.Greeting
	if err := r.db.Where("message = ?", message).First(&greeting).Error; err != nil {
		return nil, err
	}
	return &greeting, nil
}

func (r *helloRepositoryImpl) ExistsByMessage(message string) (bool, error) {
	var count int64
	if err := r.db.Model(&entity.Greeting{}).Where("message = ?", message).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
