package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Function to create hashed password using bcrypt package
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}

	return string(bytes), err
}

// Function to verify saved hashed password with input password
func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
