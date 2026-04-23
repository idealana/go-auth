package utils

import (
	"time"
	"crypto/rand"
	"encoding/hex"
	
	"github.com/golang-jwt/jwt/v5"
	"go-auth/internal/model/domain"
)

type TokenUtility struct {
	AccessKey string
	AccessExpired int
}

func NewTokenUtility(accessKey string, accessExpired int) *TokenUtility {
	return &TokenUtility{
		AccessKey: accessKey,
		AccessExpired: accessExpired,
	}
}

func (t *TokenUtility) generateJWT(user *domain.User, key string, expired int) (string, error) {
	claims := jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Minute * time.Duration(expired)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(key))
}

func (t *TokenUtility) GenerateAccessToken(user *domain.User) (string, error) {
	return t.generateJWT(user, t.AccessKey, t.AccessExpired)
}

func (t *TokenUtility) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func (t *JWT) GenerateToken(user *domain.User) (string, string, error) {
	accessToken, err := t.GenerateAccessToken(user)
	if err != nil {
		return "", "", err
	}
	
	refreshToken, err := t.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
