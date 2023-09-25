package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)

	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GenerateRandomString(n int) (string, error) {
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
