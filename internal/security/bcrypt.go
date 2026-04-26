package security

import (
	"errors"
	"fmt"
	
	"golang.org/x/crypto/bcrypt"
)

var (
    ErrInvalidPassword = errors.New("invalid password")
    ErrEmptyPassword = errors.New("password cannot be empty")
)

type PasswordManager interface {
	Verify(password, hash string) error
	Hash(password string) (string, error)
}

type BcryptPassword struct {
	cost int
}

func NewBcryptPassword(cost int) PasswordManager {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
        cost = bcrypt.DefaultCost
    }
	
	return &BcryptPassword{
		cost: cost,
	}
}

func (p *BcryptPassword) Verify(password, hash string) error {
	if password == "" {
		return ErrEmptyPassword
	}
	
	if hash == "" {
		return ErrInvalidPassword
	}
	
    if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return ErrInvalidPassword
	}
	
	return nil
}

func (p *BcryptPassword) Hash(password string) (string, error) {
	if password == "" {
		return "", ErrEmptyPassword
	}
	
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}
	
    return string(hashed), nil
}
