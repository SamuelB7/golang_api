package auth

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateJwtToken(secret []byte, userID string) (string, error) {
	err := godotenv.Load()
	// err := godotenv.Load("../../.env") for testing
	if err != nil {
		log.Fatal("Error loading .env file on auth/jwt.go")
	}

	expirationSeconds, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_SECONDS"))
	if err != nil {
		log.Fatal("Error converting JWT_EXPIRATION_SECONDS to int")
	}

	expiration := time.Second * time.Duration(expirationSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"expireAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
