package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"go-auth/internal/model/domain"
)

type JWT struct {
	AccessKey string
	AccessExpired int
}

func NewJWT(accessKey string, accessExpired int) *JWT {
	return &JWT{
		AccessKey: accessKey,
		AccessExpired: accessExpired,
	}
}

func (t JWT) GenerateToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"expire": time.Now().Add(time.Minute * time.Duration(t.AccessExpired)).UnixMilli(),
	})

	jwtToken, err := token.SignedString([]byte(t.AccessKey))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
