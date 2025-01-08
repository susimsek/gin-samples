package config

import (
	"gin-samples/internal/util"
	"log"
	"path/filepath"
)

// JweTokenInitializer interface for initializing JWE Key Pairs
type JweTokenInitializer interface {
	InitJweKeyPair() (*util.RSAKeyPair, *util.RSAKeyPair)
}

// RealJweTokenConfig is the production implementation
type RealJweTokenConfig struct{}

// JweTokenConfig is the default implementation for production
var JweTokenConfig JweTokenInitializer = &RealJweTokenConfig{}

// InitJweKeyPair loads RSA key pairs for signing and encryption from files
func (r *RealJweTokenConfig) InitJweKeyPair() (*util.RSAKeyPair, *util.RSAKeyPair) {
	// Paths for signing keys
	signPrivateKeyPath := filepath.Join("resources", "keys", "sign", "private_key.pem")
	signPublicKeyPath := filepath.Join("resources", "keys", "sign", "public_key.pem")

	// Paths for encryption keys
	encPrivateKeyPath := filepath.Join("resources", "keys", "enc", "private_key.pem")
	encPublicKeyPath := filepath.Join("resources", "keys", "enc", "public_key.pem")

	// Load signing key pair
	signKeyPair, err := util.LoadRSAKeyPair(signPrivateKeyPath, signPublicKeyPath)
	if err != nil {
		log.Fatalf("Failed to load signing RSA key pair: %v", err)
	}

	// Load encryption key pair
	encKeyPair, err := util.LoadRSAKeyPair(encPrivateKeyPath, encPublicKeyPath)
	if err != nil {
		log.Fatalf("Failed to load encryption RSA key pair: %v", err)
	}

	log.Println("JWE Key Pairs loaded successfully!")
	return signKeyPair, encKeyPair
}
