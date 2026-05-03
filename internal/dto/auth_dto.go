package dto

type LoginRequest struct {
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,max=20"`
}

type LoginResponse struct {
    UserID int `json:"user_id"`
    AccessToken string `json:"access_token"`
}

type LoginResult struct {
    UserID int
    AccessToken string
    RefreshToken string
}
