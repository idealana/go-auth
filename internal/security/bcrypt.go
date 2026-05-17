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

func NewBcryptPassword(cost ...int) PasswordManager {
	selectedCost := bcrypt.DefaultCost

	if len(cost) > 0 &&
		cost[0] >= bcrypt.MinCost &&
		cost[0] <= bcrypt.MaxCost {
		selectedCost = cost[0]
	}
	
	return &BcryptPassword{
		cost: selectedCost,
	}
}

func (p *BcryptPassword) Verify(hashed, plain string) error {
	if hashed == "" {
		return ErrInvalidPassword
	}

	if plain == "" {
		return ErrEmptyPassword
	}
	
    if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)); err != nil {
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
