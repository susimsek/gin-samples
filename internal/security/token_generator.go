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
	Issuer      string   `json:"iss"`
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
	signKeyPair   *util.RSAKeyPair
	encKeyPair    *util.RSAKeyPair
	tokenDuration time.Duration
	issuer        string
}

// NewTokenGenerator creates a new instance of TokenGenerator
func NewTokenGenerator(signKeyPair, encKeyPair *util.RSAKeyPair, tokenDuration time.Duration, issuer string) TokenGenerator {
	return &tokenGenerator{
		signKeyPair:   signKeyPair,
		encKeyPair:    encKeyPair,
		tokenDuration: tokenDuration,
		issuer:        issuer,
	}
}

// Generate creates a signed and encrypted JWE token
func (t *tokenGenerator) Generate(claims TokenClaims) (Token, error) {
	now := time.Now().Unix()
	claims = t.populateStandardClaims(claims, now)

	claimsBytes, err := t.serializeClaims(claims)
	if err != nil {
		return Token{}, err
	}

	signedPayload, err := t.signClaims(claimsBytes)
	if err != nil {
		return Token{}, err
	}

	encryptedPayload, err := t.encryptPayload(signedPayload)
	if err != nil {
		return Token{}, err
	}

	return Token{
		AccessToken: encryptedPayload,
		TokenType:   "Bearer",
		ExpiresIn:   int64(t.tokenDuration.Seconds()),
	}, nil
}

// Validate validates and decrypts a JWE token to extract TokenClaims
func (t *tokenGenerator) Validate(tokenString string) (*TokenClaims, error) {
	decryptedBytes, err := t.decryptToken(tokenString)
	if err != nil {
		return nil, err
	}

	verifiedBytes, err := t.verifySignature(decryptedBytes)
	if err != nil {
		return nil, err
	}

	claims, err := t.deserializeClaims(verifiedBytes)
	if err != nil {
		return nil, err
	}

	if err := t.validateTokenClaims(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// Private Methods

func (t *tokenGenerator) populateStandardClaims(claims TokenClaims, now int64) TokenClaims {
	claims.IssuedAt = now
	claims.ExpiresAt = now + int64(t.tokenDuration.Seconds())
	claims.NotBefore = now
	claims.JTI = uuid.NewString()
	claims.Issuer = t.issuer
	return claims
}

func (t *tokenGenerator) serializeClaims(claims TokenClaims) ([]byte, error) {
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return nil, errors.New("failed to serialize claims: " + err.Error())
	}
	return claimsBytes, nil
}

func (t *tokenGenerator) signClaims(claimsBytes []byte) (string, error) {
	signingKey := jose.SigningKey{Algorithm: jose.RS256, Key: t.signKeyPair.PrivateKey}
	signer, err := jose.NewSigner(signingKey, nil)
	if err != nil {
		return "", errors.New("failed to create signer: " + err.Error())
	}

	signedObject, err := signer.Sign(claimsBytes)
	if err != nil {
		return "", errors.New("failed to sign claims: " + err.Error())
	}

	return signedObject.CompactSerialize()
}

func (t *tokenGenerator) encryptPayload(signedPayload string) (string, error) {
	encrypter, err := jose.NewEncrypter(
		jose.A256GCM,
		jose.Recipient{Algorithm: jose.RSA_OAEP_256, Key: t.encKeyPair.PublicKey},
		nil,
	)
	if err != nil {
		return "", errors.New("failed to create encrypter: " + err.Error())
	}

	encryptedObject, err := encrypter.Encrypt([]byte(signedPayload))
	if err != nil {
		return "", errors.New("failed to encrypt payload: " + err.Error())
	}

	return encryptedObject.CompactSerialize()
}

func (t *tokenGenerator) decryptToken(tokenString string) ([]byte, error) {
	keyAlgorithms := []jose.KeyAlgorithm{
		jose.RSA_OAEP_256,
	}
	contentEncryption := []jose.ContentEncryption{
		jose.A256GCM,
	}
	object, err := jose.ParseEncrypted(tokenString, keyAlgorithms, contentEncryption)

	if err != nil {
		return nil, &customError.JwtError{Message: "Failed to parse JWE: " + err.Error()}
	}

	return object.Decrypt(t.encKeyPair.PrivateKey)
}

func (t *tokenGenerator) verifySignature(decryptedBytes []byte) ([]byte, error) {
	signatureAlgorithms := []jose.SignatureAlgorithm{
		jose.RS256,
	}
	signedObject, err := jose.ParseSigned(string(decryptedBytes), signatureAlgorithms)
	if err != nil {
		return nil, &customError.JwtError{Message: "Failed to parse signed payload: " + err.Error()}
	}

	return signedObject.Verify(t.signKeyPair.PublicKey)
}

func (t *tokenGenerator) deserializeClaims(verifiedBytes []byte) (*TokenClaims, error) {
	var claims TokenClaims
	err := json.Unmarshal(verifiedBytes, &claims)
	if err != nil {
		return nil, &customError.JwtError{Message: "Failed to deserialize claims: " + err.Error()}
	}
	return &claims, nil
}

func (t *tokenGenerator) validateTokenClaims(claims *TokenClaims) error {
	now := time.Now().Unix()
	if now > claims.ExpiresAt {
		return &customError.JwtError{Message: "Token has expired"}
	}
	if now < claims.NotBefore {
		return &customError.JwtError{Message: "Token is not valid yet"}
	}
	return nil
}
