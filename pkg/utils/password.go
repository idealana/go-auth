package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(hash, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func HashPassword(password string, cost int) (string, error) {
    if cost == 0 {
        cost = 12
    }
    
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    return string(bytes), err
}
