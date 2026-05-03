package security

import (
	"errors"
	"time"
	"fmt"
	"crypto/rand"
	"encoding/hex"
	
	"github.com/golang-jwt/jwt/v5"
)

var ErrGenerateToken = errors.New("failed to generate token")

type JWTGenerator interface {
    GenerateAccessToken(user UserClaims) (string, error)
    GenerateToken(user UserClaims) (TokenPair, error)
}

type RefreshTokenGenerator interface {
    GenerateRefreshToken() (string, error)
}

type UserClaims interface {
    GetUserID() int
}

type TokenGenerator interface {
	JWTGenerator
	RefreshTokenGenerator
}

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTAuthToken struct {
	appName string
	accessKey []byte
	accessExpired time.Duration
	signingMethod jwt.SigningMethod
}

type TokenPair struct {
    AccessToken  string
    RefreshToken string
}

func NewJWTAuthToken(appName, accessKey string, accessExpired time.Duration) (TokenGenerator, error) {
	if appName == "" {
        return nil, errors.New("app name is required")
    }
    if accessKey == "" {
        return nil, errors.New("access key is required")
    }
    if accessExpired <= 0 {
        return nil, errors.New("access expired must be positive")
    }
	
	return &JWTAuthToken{
		appName: appName,
		accessKey: []byte(accessKey),
		accessExpired: accessExpired,
		signingMethod:  jwt.SigningMethodHS256, // default
	}, nil
}

func (t *JWTAuthToken) generateJWT(user UserClaims, key []byte, expired time.Duration) (string, error) {
	if user == nil {
		return "", errors.New("generate jwt: user is required")
	}

	now := time.Now()

	claims := JWTClaims{
		UserID: user.GetUserID(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expired)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    t.appName,
		},
	}

	token := jwt.NewWithClaims(t.signingMethod, claims)
	
	return token.SignedString(key)
}

func (t *JWTAuthToken) GenerateAccessToken(user UserClaims) (string, error) {
	return t.generateJWT(user, t.accessKey, t.accessExpired)
}

func (t *JWTAuthToken) GenerateRefreshToken() (string, error) {
	const byteSize = 32
	b := make([]byte, byteSize)
	
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate refresh token: %w", err)
	}

	return hex.EncodeToString(b), nil
}

func (t *JWTAuthToken) GenerateToken(user UserClaims) (TokenPair, error) {
	accessToken, err := t.GenerateAccessToken(user)
	if err != nil {
		return TokenPair{}, fmt.Errorf("%w: %w", ErrGenerateToken, err)
	}
	
	refreshToken, err := t.GenerateRefreshToken()
	if err != nil {
		return TokenPair{}, fmt.Errorf("%w: %w", ErrGenerateToken, err)
	}

	return TokenPair{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}
