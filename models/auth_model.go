package models

import "github.com/golang-jwt/jwt"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type JWTClaims struct {
	*jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"user_id"`
}

type ErrorResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// DB related
type FindByUsernameResponse struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
