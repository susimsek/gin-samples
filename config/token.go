package config

import (
	"gin-samples/internal/util"
	"log"
	"path/filepath"
)

// TokenInitializer interface for initializing JWT Key Pair
type TokenInitializer interface {
	InitJwtKeyPair() *util.RSAKeyPair
}

// RealTokenConfig is the production implementation
type RealTokenConfig struct{}

// TokenConfig is the default implementation for production
var TokenConfig TokenInitializer = &RealTokenConfig{}

// InitJwtKeyPair for RealTokenConfig loads RSA key pair from files
func (r *RealTokenConfig) InitJwtKeyPair() *util.RSAKeyPair {
	privateKeyPath := filepath.Join("resources", "keys", "private_key.pem")
	publicKeyPath := filepath.Join("resources", "keys", "public_key.pem")

	jwtKeyPair, err := util.LoadRSAKeyPair(privateKeyPath, publicKeyPath)
	if err != nil {
		log.Fatalf("Failed to load RSA key pair: %v", err)
	}

	log.Println("RSA Key Pair loaded successfully!")
	return jwtKeyPair
}
