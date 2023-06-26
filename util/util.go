package util

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))

	hashedPassword := fmt.Sprintf("%x", hash)
	return hashedPassword
}
func VerifyPassword(password, hashedPassword string) bool {
	hashedInput := HashPassword(password)

	if hashedInput == hashedPassword {
		return true
	}
	return false
}
