package dto

type LoginRequest struct {
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,max=20"`
}

type LoginResponse struct {
    UserID uint `json:"user_id"`
    AccessToken string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

type LoginResult struct {
    UserID uint
    AccessToken string
    RefreshToken string
}

type ProfileResponse struct {
    UserID uint `json:"user_id"`
    Email string `json:"email"`
    Role string `json:"role"`
    Status string `json:"status"`
}

type ProfileResult struct {
    UserID uint
    Email string
    Role string
    Status string
}
