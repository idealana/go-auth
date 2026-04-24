package utils

import (
	"errors"
	
	"golang.org/x/crypto/bcrypt"
)

var (
    ErrInvalidPassword = errors.New("invalid password")
    ErrEmptyPassword = errors.New("password cannot be empty")
)

type Password struct {
	cost int
}

func NewPassword(cost int) *Password {
	if cost == 0 {
        cost = bcrypt.DefaultCost
    }
	
	return &Password{
		cost: cost,
	}
}

func (p *Password) Verify(hash, password string) error {
	if hash == "" {
		return ErrInvalidPassword
	}

	if password == "" {
		return ErrEmptyPassword
	}
	
    if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return ErrInvalidPassword
	}
	
	return nil
}

func (p *Password) Hash(password string) (string, error) {
	if password == "" {
		return "", ErrEmptyPassword
	}
	
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	if err != nil {
		return "", err
	}
	
    return string(bytes), nil
}
