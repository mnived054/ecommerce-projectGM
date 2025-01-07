package services

import (
	"ecommerce-project/backend/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretkey = []byte("supersecretkey")

func RegisterUser(user models.User) error {

	return nil
}

func LoginUser(credentials models.LoginCredentials) (string, error) {

	if credentials.Username != "admin" || credentials.Password != "admin" {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": credentials.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
