package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the provided password
// returns the hashed password and an error if it exists
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %s", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
// returns nil if the password is correct and an error if it is not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
