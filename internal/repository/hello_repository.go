package repository

import (
	"gin-samples/internal/model"
	"sync"
)

type HelloRepository interface {
	SaveGreeting(input model.GreetingInput) model.Greeting
	GetAllGreetings() []model.Greeting
	FindByMessage(message string) (*model.Greeting, bool)
	ExistsByMessage(message string) bool
}

type helloRepositoryImpl struct {
	data  []model.Greeting
	mutex sync.Mutex
}

func NewHelloRepository() HelloRepository {
	return &helloRepositoryImpl{
		data: make([]model.Greeting, 0),
	}
}

func (r *helloRepositoryImpl) SaveGreeting(input model.GreetingInput) model.Greeting {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	newGreeting := model.Greeting{
		Message: input.Message,
	}
	r.data = append(r.data, newGreeting)

	return newGreeting
}

func (r *helloRepositoryImpl) GetAllGreetings() []model.Greeting {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.data
}

func (r *helloRepositoryImpl) FindByMessage(message string) (*model.Greeting, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, greeting := range r.data {
		if greeting.Message == message {
			return &greeting, true
		}
	}
	return nil, false
}

func (r *helloRepositoryImpl) ExistsByMessage(message string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, greeting := range r.data {
		if greeting.Message == message {
			return true
		}
	}
	return false
}
