package auth

import "testing"

func TestCreateJwt(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJwtToken(secret, "123")
	if err != nil {
		t.Errorf("Error creating token: %v", err)
	}

	if token == "" {
		t.Errorf("Expected a token, got an empty string")
	}
}
