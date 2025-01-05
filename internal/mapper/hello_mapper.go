package mapper

import (
	"gin-samples/internal/domain"
	"gin-samples/internal/dto"
)

// HelloMapper defines the interface for mapping operations related to greetings
type HelloMapper interface {
	ToGreetingResponse(domain.Greeting) (dto.GreetingResponse, error)
	ToGreetingResponses([]domain.Greeting) ([]dto.GreetingResponse, error)
	ToGreetingEntity(dto.GreetingInput) (domain.Greeting, error)
}

// helloMapperImpl is the default implementation of HelloMapper
type helloMapperImpl struct{}

// NewHelloMapper creates a new instance of helloMapperImpl
func NewHelloMapper() HelloMapper {
	return &helloMapperImpl{}
}

// ToGreetingResponse maps a Greeting domain to GreetingResponse DTO
func (m *helloMapperImpl) ToGreetingResponse(g domain.Greeting) (dto.GreetingResponse, error) {
	return dto.GreetingResponse{
		ID:        g.ID,
		Message:   g.Message,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}

// ToGreetingResponses maps a slice of Greeting entities to GreetingResponse DTOs
func (m *helloMapperImpl) ToGreetingResponses(greetings []domain.Greeting) ([]dto.GreetingResponse, error) {
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

// ToGreetingEntity maps a GreetingInput DTO to a Greeting domain
func (m *helloMapperImpl) ToGreetingEntity(input dto.GreetingInput) (domain.Greeting, error) {
	return domain.Greeting{
		Message: input.Message,
	}, nil
}
