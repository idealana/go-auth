package utils

import (
	"errors"
	"time"
	"fmt"
	"crypto/rand"
	"encoding/hex"
	
	"github.com/golang-jwt/jwt/v5"
	"go-auth/internal/model/domain"
)

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenUtility struct {
	AppName string
	AccessKey string
	AccessExpired time.Duration
}

func NewTokenUtility(appName, accessKey string, accessExpired time.Duration) *TokenUtility {
	return &TokenUtility{
		AppName: appName,
		AccessKey: accessKey,
		AccessExpired: accessExpired,
	}
}

func (t *TokenUtility) generateJWT(user *domain.User, key string, expired int) (string, error) {
	if user == nil {
		return "", errors.New("User not found.")
	}

	exp := time.Now().Add(t.AccessExpired)
	now := time.Now()

	claims := JWTClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    t.AppName,
		},
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

func (t *TokenUtility) GenerateToken(user *domain.User) (string, string, error) {
	accessToken, err := t.GenerateAccessToken(user)
	if err != nil {
		return "", "", fmt.Errorf("generate access token: %w", err)
	}
	
	refreshToken, err := t.GenerateRefreshToken()
	if err != nil {
		return "", "", fmt.Errorf("generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}
