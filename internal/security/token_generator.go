package security

import (
	"encoding/json"
	"errors"
	customError "gin-samples/internal/error"
	"gin-samples/internal/util"
	"github.com/go-jose/go-jose/v4"
	"github.com/google/uuid"
	"time"
)

// TokenClaims represents claims for the JWE token
type TokenClaims struct {
	UserID      string   `json:"sub"`
	Authorities []string `json:"authorities"`
	IssuedAt    int64    `json:"iat"`
	ExpiresAt   int64    `json:"exp"`
	NotBefore   int64    `json:"nbf"`
	JTI         string   `json:"jti"`
}

// Token represents the JWE token structure
type Token struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
}

// TokenGenerator defines the interface for generating and validating tokens
type TokenGenerator interface {
	Generate(claims TokenClaims) (Token, error)
	Validate(tokenString string) (*TokenClaims, error)
}

type tokenGenerator struct {
	signKeyPair   *util.RSAKeyPair // Signing key pair
	encKeyPair    *util.RSAKeyPair // Encryption key pair
	tokenDuration time.Duration
}

// NewTokenGenerator creates a new instance of TokenGenerator
func NewTokenGenerator(signKeyPair, encKeyPair *util.RSAKeyPair, tokenDuration time.Duration) TokenGenerator {
	return &tokenGenerator{
		signKeyPair:   signKeyPair,
		encKeyPair:    encKeyPair,
		tokenDuration: tokenDuration,
	}
}

// Generate creates a signed and encrypted JWE token
func (t *tokenGenerator) Generate(claims TokenClaims) (Token, error) {
	now := time.Now().Unix()
	expiration := now + int64(t.tokenDuration.Seconds())

	// Populate standard claims
	claims.IssuedAt = now
	claims.ExpiresAt = expiration
	claims.NotBefore = now
	claims.JTI = uuid.NewString()

	// Serialize claims to JSON
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return Token{}, errors.New("failed to serialize claims: " + err.Error())
	}

	// Sign the claims
	signingKey := jose.SigningKey{Algorithm: jose.RS256, Key: t.signKeyPair.PrivateKey}
	signer, err := jose.NewSigner(signingKey, nil)
	if err != nil {
		return Token{}, errors.New("failed to create signer: " + err.Error())
	}

	signedObject, err := signer.Sign(claimsBytes)
	if err != nil {
		return Token{}, errors.New("failed to sign claims: " + err.Error())
	}

	// Serialize signed payload
	signedPayload, err := signedObject.CompactSerialize()
	if err != nil {
		return Token{}, errors.New("failed to serialize signed payload: " + err.Error())
	}

	// Encrypt the signed payload
	encrypter, err := jose.NewEncrypter(
		jose.A256GCM,
		jose.Recipient{Algorithm: jose.RSA_OAEP_256, Key: t.encKeyPair.PublicKey},
		nil,
	)
	if err != nil {
		return Token{}, errors.New("failed to create encrypter: " + err.Error())
	}

	encryptedObject, err := encrypter.Encrypt([]byte(signedPayload))
	if err != nil {
		return Token{}, errors.New("failed to encrypt payload: " + err.Error())
	}

	encryptedPayload, err := encryptedObject.CompactSerialize()
	if err != nil {
		return Token{}, errors.New("failed to serialize encrypted payload: " + err.Error())
	}

	return Token{
		AccessToken: encryptedPayload,
		TokenType:   "Bearer",
		ExpiresIn:   int64(t.tokenDuration.Seconds()),
	}, nil
}

// Validate validates and decrypts a JWE token to extract TokenClaims
func (t *tokenGenerator) Validate(tokenString string) (*TokenClaims, error) {
	// Parse the encrypted token
	keyAlgorithms := []jose.KeyAlgorithm{
		jose.RSA_OAEP_256,
	}
	contentEncryption := []jose.ContentEncryption{
		jose.A256GCM,
	}

	// Parse the encrypted token
	object, err := jose.ParseEncrypted(tokenString, keyAlgorithms, contentEncryption)
	if err != nil {
		return nil, &customError.JwtError{Message: "Failed to parse JWE: " + err.Error()}
	}

	// Decrypt the token using the encryption key pair's private key
	decryptedBytes, err := object.Decrypt(t.encKeyPair.PrivateKey)
	if err != nil {
		return nil, &customError.JwtError{Message: "Failed to decrypt JWE: " + err.Error()}
	}

	signatureAlgorithms := []jose.SignatureAlgorithm{
		jose.RS256,
	}

	// Parse the signed payload
	signedObject, err := jose.ParseSigned(string(decryptedBytes), signatureAlgorithms)
	if err != nil {
		return nil, &customError.JwtError{Message: "Failed to parse signed payload: " + err.Error()}
	}

	// Verify the signature
	verifiedBytes, err := signedObject.Verify(t.signKeyPair.PublicKey)
	if err != nil {
		return nil, &customError.JwtError{Message: "Signature verification failed: " + err.Error()}
	}

	// Deserialize claims
	var claims TokenClaims
	err = json.Unmarshal(verifiedBytes, &claims)
	if err != nil {
		return nil, &customError.JwtError{Message: "Failed to deserialize claims: " + err.Error()}
	}

	// Validate expiration
	now := time.Now().Unix()
	if now > claims.ExpiresAt {
		return nil, &customError.JwtError{Message: "Token has expired"}
	}

	// Validate "not before" claim
	if now < claims.NotBefore {
		return nil, &customError.JwtError{Message: "Token is not valid yet"}
	}

	return &claims, nil
}
