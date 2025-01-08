package mock

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"gin-samples/internal/util"
)

// MockJweTokenConfig implements the JweTokenInitializer interface for testing purposes
type MockJweTokenConfig struct{}

// InitJweKeyPair dynamically generates mock RSA key pairs for signing and encryption
func (m *MockJweTokenConfig) InitJweKeyPair() (*util.RSAKeyPair, *util.RSAKeyPair) {
	// Generate signing key pair
	signPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate mock signing RSA private key: %v", err)
	}
	signPublicKey := &signPrivateKey.PublicKey
	signKeyPair := &util.RSAKeyPair{
		PrivateKey: signPrivateKey,
		PublicKey:  signPublicKey,
	}

	// Generate encryption key pair
	encPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate mock encryption RSA private key: %v", err)
	}
	encPublicKey := &encPrivateKey.PublicKey
	encKeyPair := &util.RSAKeyPair{
		PrivateKey: encPrivateKey,
		PublicKey:  encPublicKey,
	}

	return signKeyPair, encKeyPair
}
