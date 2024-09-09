package auth

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateJwt(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file for testing")
	}

	secret := []byte("secret")

	token, err := CreateJwtToken(secret, "123")
	if err != nil {
		t.Errorf("Error creating token: %v", err)
	}

	if token == "" {
		t.Errorf("Expected a token, got an empty string")
	}
}
