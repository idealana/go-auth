package dto

type LoginRequest struct {
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,max=20"`
}

type LoginResponse struct {
    UserID uint `json:"user_id"`
    AccessToken string `json:"access_token"`
}

type LoginResult struct {
    UserID uint
    AccessToken string
    RefreshToken string
}
