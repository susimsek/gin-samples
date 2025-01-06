package mapper

import (
	"gin-samples/internal/domain"
	"gin-samples/internal/dto"
)

// HelloMapper defines the interface for mapping operations related to greetings
type HelloMapper interface {
	ToGreetingResponse(domain.Greeting) dto.GreetingResponse
	ToGreetingResponses([]domain.Greeting) []dto.GreetingResponse
	ToGreetingEntity(dto.GreetingInput) domain.Greeting
	PartialUpdateGreeting(*domain.Greeting, dto.GreetingInput)
}

// helloMapperImpl is the default implementation of HelloMapper
type helloMapperImpl struct{}

// NewHelloMapper creates a new instance of helloMapperImpl
func NewHelloMapper() HelloMapper {
	return &helloMapperImpl{}
}

// ToGreetingResponse maps a Greeting domain to GreetingResponse DTO
func (m *helloMapperImpl) ToGreetingResponse(g domain.Greeting) dto.GreetingResponse {
	return dto.GreetingResponse{
		ID:        g.ID,
		Message:   g.Message,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

// ToGreetingResponses maps a slice of Greeting entities to GreetingResponse DTOs
func (m *helloMapperImpl) ToGreetingResponses(greetings []domain.Greeting) []dto.GreetingResponse {
	responses := make([]dto.GreetingResponse, len(greetings))
	for i, g := range greetings {
		responses[i] = m.ToGreetingResponse(g)
	}
	return responses
}

// ToGreetingEntity maps a GreetingInput DTO to a Greeting domain
func (m *helloMapperImpl) ToGreetingEntity(input dto.GreetingInput) domain.Greeting {
	return domain.Greeting{
		Message: input.Message,
	}
}

// PartialUpdateGreeting updates only the provided fields in the domain entity
func (m *helloMapperImpl) PartialUpdateGreeting(entity *domain.Greeting, input dto.GreetingInput) {
	if input.Message != "" {
		entity.Message = input.Message
	}
}
