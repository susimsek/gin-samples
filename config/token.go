package config

import (
	"gin-samples/internal/util"
	"log"
	"path/filepath"
)

// InitJwtKeyPair initializes the RSA key pair for JWT
func InitJwtKeyPair() *util.RSAKeyPair {
	privateKeyPath := filepath.Join("resources", "keys", "private_key.pem")
	publicKeyPath := filepath.Join("resources", "keys", "public_key.pem")

	jwtKeyPair, err := util.LoadRSAKeyPair(privateKeyPath, publicKeyPath)
	if err != nil {
		log.Fatalf("Failed to load RSA key pair: %v", err)
	}

	return jwtKeyPair
}
