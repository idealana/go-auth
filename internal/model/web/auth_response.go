package web

import (
	"go-auth/internal/model/domain"
)

type LoginResponse struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
}

func NewLoginResponse(user domain.User) LoginResponse {
	return LoginResponse{
		FullName: user.FullName,
		Email: user.Email,
	}
}
