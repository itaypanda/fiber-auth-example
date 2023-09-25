// https://stackoverflow.com/questions/23259586/bcrypt-password-hashing-in-golang-compatible-with-node-js

package utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GetCost() int {
	cost, err := strconv.Atoi(os.Getenv("BCRYPT_ROUNDS"))
	if err != nil {
		cost = 12
	}

	return cost
}

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, GetCost())
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CompareHashedPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
