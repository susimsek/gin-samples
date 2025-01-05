package mapper

import (
	"gin-samples/internal/dto"
	"gin-samples/internal/entity"
)

// HelloMapper defines the interface for mapping operations related to greetings
type HelloMapper interface {
	ToGreetingResponse(entity.Greeting) (dto.GreetingResponse, error)
	ToGreetingResponses([]entity.Greeting) ([]dto.GreetingResponse, error)
	ToGreetingEntity(dto.GreetingInput) (entity.Greeting, error)
}

// helloMapperImpl is the default implementation of HelloMapper
type helloMapperImpl struct{}

// NewHelloMapper creates a new instance of helloMapperImpl
func NewHelloMapper() HelloMapper {
	return &helloMapperImpl{}
}

// ToGreetingResponse maps a Greeting entity to GreetingResponse DTO
func (m *helloMapperImpl) ToGreetingResponse(g entity.Greeting) (dto.GreetingResponse, error) {
	return dto.GreetingResponse{
		ID:        g.ID,
		Message:   g.Message,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}

// ToGreetingResponses maps a slice of Greeting entities to GreetingResponse DTOs
func (m *helloMapperImpl) ToGreetingResponses(greetings []entity.Greeting) ([]dto.GreetingResponse, error) {
	responses := make([]dto.GreetingResponse, len(greetings))
	for i, g := range greetings {
		responses[i] = dto.GreetingResponse{
			ID:        g.ID,
			Message:   g.Message,
			CreatedAt: g.CreatedAt,
			UpdatedAt: g.UpdatedAt,
		}
	}
	return responses, nil
}

// ToGreetingEntity maps a GreetingInput DTO to a Greeting entity
func (m *helloMapperImpl) ToGreetingEntity(input dto.GreetingInput) (entity.Greeting, error) {
	return entity.Greeting{
		Message: input.Message,
	}, nil
}
