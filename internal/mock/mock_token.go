package mock

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"gin-samples/internal/util"
)

// MockTokenConfig implements the TokenInitializer interface for testing purposes
type MockTokenConfig struct{}

// InitJwtKeyPair dynamically generates a mock RSA key pair for testing
func (m *MockTokenConfig) InitJwtKeyPair() *util.RSAKeyPair {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate mock RSA private key: %v", err)
	}

	publicKey := &privateKey.PublicKey

	return &util.RSAKeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}
