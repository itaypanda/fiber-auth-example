package utils

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/itaypanda/fiber-auth-example/models"
)

func GetJWTSecret() ([]byte, error) {
	var err error

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Print("JWT secret was not found, generating debug one (WILL NOT PERSIST, FOR DEVELOPMENT ONLY!)")
		secret, err = GenerateRandomString(32)
		if err != nil {
			return nil, err
		}
	}

	return []byte(secret), nil
}

func GetJWTLifespan() int {
	cost, err := strconv.Atoi(os.Getenv("JWT_LIFESPAN"))
	if err != nil {
		cost = 24
	}

	return cost
}

func UserSignedClaims(user *models.User) (string, error) {
	var err error

	claims := jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(GetJWTLifespan())).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret, err := GetJWTSecret()
	if err != nil {
		return "", err
	}
	signedToken, err := token.SignedString(secret)
	if err != nil {
			return "", err
	}

	return signedToken, nil
}
