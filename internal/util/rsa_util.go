package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// RSAKeyPair represents a pair of RSA keys (private and public).
type RSAKeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// LoadRSAKeyPair loads the RSA private and public keys from given file paths.
func LoadRSAKeyPair(privateKeyPath, publicKeyPath string) (*RSAKeyPair, error) {
	privateKey, err := loadRSAPrivateKey(privateKeyPath)
	if err != nil {
		return nil, err
	}

	publicKey, err := loadRSAPublicKey(publicKeyPath)
	if err != nil {
		return nil, err
	}

	return &RSAKeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

// loadRSAPrivateKey loads the RSA private key from a PEM file.
func loadRSAPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(fileContent)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key: unsupported format")
	}

	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("parsed key is not an RSA private key")
	}

	return privateKey, nil
}

// loadRSAPublicKey loads the RSA public key from a PEM file.
func loadRSAPublicKey(filePath string) (*rsa.PublicKey, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(fileContent)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse public key: unsupported format")
	}

	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("parsed key is not an RSA public key")
	}

	return publicKey, nil
}
