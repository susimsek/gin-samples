package dto

// TokenResponse represents the JWT token response
// @Description JWT token response DTO
type TokenResponse struct {
	// AccessToken is the JWT access token
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." validate:"required"`

	// TokenType is the type of the token
	TokenType string `json:"tokenType" example:"Bearer" validate:"required"`

	// AccessTokenExpiresIn is the expiration time of the access token in seconds
	AccessTokenExpiresIn int64 `json:"accessTokenExpiresIn" example:"3600" validate:"required"`
}
