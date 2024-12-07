package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var validate *validator.Validate
var db *gorm.DB
var errs error

var user Users
var register RegisterUser

type RegisterUser struct {
	Name                  string `json:"name" validate:"required,min=5,max=30"`
	Email                 string `json:"email" validate:"required,email"`
	Password              string `json:"password" validate:"required,min=8,max=30"`
	Password_Confirmation string `json:"password_confirmation" validate:"required,min=8,max=30,eqfield=Password"`
}

type Users struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

type JWTClaim struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, username string) (string, error) {
	privateKeyBytes, err := os.ReadFile("./private/private.pem")
	if err != nil {
		return "", fmt.Errorf("could not read private key: %v", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("could not parse private key: %v", err)
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		Name:  username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %v", err)
	}
	return tokenString, err
}
