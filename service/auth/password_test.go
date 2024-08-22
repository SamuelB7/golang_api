package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("Expected a hash, got an empty string")
	}

	if hash == "password" {
		t.Errorf("Expected a hash, got the original password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePassword(hash, "password") {
		t.Errorf("Expected the passwords to match")
	}

	if ComparePassword(hash, "wrongpassword") {
		t.Errorf("Expected the passwords not to match")
	}
}
