package security

import (
	"errors"
	customError "gin-samples/internal/error"
	"gin-samples/internal/util"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// TokenClaims represents claims for the JWT token
type TokenClaims struct {
	UserID      string
	Authorities []string
}

// Token represents the JWT token structure
type Token struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
}

// TokenGenerator defines the interface for token generation and validation
type TokenGenerator interface {
	Generate(claims TokenClaims) (Token, error)
	Validate(tokenString string) (*jwt.MapClaims, error)
}

type jwtTokenGenerator struct {
	keyPair       *util.RSAKeyPair // Pointer to RSAKeyPair
	tokenDuration time.Duration
}

// NewTokenGenerator creates a new instance of TokenGenerator
func NewTokenGenerator(keyPair *util.RSAKeyPair, tokenDuration time.Duration) TokenGenerator {
	return &jwtTokenGenerator{
		keyPair:       keyPair,
		tokenDuration: tokenDuration,
	}
}

// Generate creates a new JWT token
func (t *jwtTokenGenerator) Generate(claims TokenClaims) (Token, error) {
	now := time.Now()
	expiration := now.Add(t.tokenDuration)

	jwtClaims := jwt.MapClaims{
		"sub":         claims.UserID,
		"authorities": claims.Authorities,
		"iat":         now.Unix(),
		"exp":         expiration.Unix(),
		"nbf":         now.Unix(),
		"jti":         uuid.NewString(),
	}

	// Sign the token with the private key
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtClaims)
	tokenString, err := token.SignedString(t.keyPair.PrivateKey)
	if err != nil {
		return Token{}, errors.New("failed to sign the token")
	}

	return Token{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   int64(t.tokenDuration.Seconds()),
	}, nil
}

// Validate validates a JWT token using the public key
func (t *jwtTokenGenerator) Validate(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, &customError.JwtError{Message: "Unexpected signing method"}
		}
		return t.keyPair.PublicKey, nil
	})

	if err != nil {
		// Handle parsing errors
		return nil, &customError.JwtError{Message: "Failed to parse token: " + err.Error()}
	}

	// Extract claims and validate the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, &customError.JwtError{Message: "Token is invalid or expired"}
	}

	return &claims, nil
}
