package services

import (
	"AuthService/models"
	"AuthService/repositories"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(requests *models.LoginRequest, repository repositories.AuthRepository) (*models.LoginResponse, error) {
	user := repository.FindByUsername(requests.Username)
	if user == nil {
		return nil, errors.New("user not found")
	}

	userResponse := models.FindByUsernameResponse{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		UserId: user.UserId,
	}

	verifyPassword := VerifyPassword(requests.Password, userResponse.Password)
	if !verifyPassword {
		return nil, errors.New("password doesn't match")
	}

	token, err := CreateToken(&userResponse)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}

func VerifyPassword(inputPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

func CreateToken(user *models.FindByUsernameResponse) (string, error) {
	claims := &models.JWTClaims{
		StandardClaims: &jwt.StandardClaims{
			Subject:   fmt.Sprint(user.UserId),
			Issuer:    "GenesysInc.",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		Username: user.Username,
		UserId: user.UserId,
	}

	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token string) (*models.JWTClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(*models.JWTClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}