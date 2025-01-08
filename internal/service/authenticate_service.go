package service

import (
	"gin-samples/internal/dto"
	customError "gin-samples/internal/error"
	"gin-samples/internal/repository"
	"gin-samples/internal/security"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticationService defines the authentication service interface
type AuthenticationService interface {
	Authenticate(input dto.LoginInput) (dto.TokenResponse, error)
}

type authenticationServiceImpl struct {
	userRepository repository.UserRepository
	tokenGenerator security.TokenGenerator
}

// NewAuthenticationService creates a new instance of AuthenticationService
func NewAuthenticationService(userRepo repository.UserRepository, tokenGen security.TokenGenerator) AuthenticationService {
	return &authenticationServiceImpl{
		userRepository: userRepo,
		tokenGenerator: tokenGen,
	}
}

// Authenticate validates the login credentials and returns a TokenResponse
func (s *authenticationServiceImpl) Authenticate(input dto.LoginInput) (dto.TokenResponse, error) {
	// Check if the user exists by username
	userOptional, err := s.userRepository.FindByUsername(input.Username)
	if err != nil {
		return dto.TokenResponse{}, err
	}

	if userOptional.IsEmpty() {
		return dto.TokenResponse{}, &customError.InvalidCredentialsError{}
	}

	user := userOptional.Value

	// Check if the user is enabled
	if !user.Enabled {
		return dto.TokenResponse{}, &customError.InvalidCredentialsError{}
	}

	// Validate the password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return dto.TokenResponse{}, &customError.InvalidCredentialsError{}
	}

	var authorities []string
	for _, roleMapping := range user.Roles {
		authorities = append(authorities, roleMapping.Role.Name)
	}

	// Generate token using TokenGenerator
	token, err := s.tokenGenerator.Generate(security.TokenClaims{
		UserID:      user.ID,
		Authorities: authorities,
	})
	if err != nil {
		return dto.TokenResponse{}, err
	}

	return dto.TokenResponse{
		AccessToken:          token.AccessToken,
		TokenType:            token.TokenType,
		AccessTokenExpiresIn: token.ExpiresIn,
	}, nil
}
