package models

type RegisterRequest struct {
	Email    string `json:"e-mail"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"e-mail"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}
