package controllers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12/middleware/jwt"
	"gorm.io/gorm"
)

var validate *validator.Validate
var db *gorm.DB
var errs error

var user Users
var register RegisterUser

const (
	accessTokenMaxAge  = 10 * time.Minute
	refreshTokenMaxAge = time.Hour
)

var (
	privateKey, publicKey = jwt.MustLoadRSA("./private/private.pem", "./private/public.pem")

	signer   = jwt.NewSigner(jwt.RS256, privateKey, accessTokenMaxAge)
	verifier = jwt.NewVerifier(jwt.RS256, publicKey)
)


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
	Name  string `json:"id"`
	Email string `json:"email"`
}

func GenerateJWT(email string, name string) (jwt.TokenPair, error) {

	refreshClaims := jwt.Claims{Subject: name}

	accessClaims := &JWTClaim{
		Name: name,
		Email: email,
	}

	tokenPair, err := signer.NewTokenPair(accessClaims, refreshClaims, refreshTokenMaxAge)
	if err != nil {
		return jwt.TokenPair{}, err
	}
	return tokenPair,nil
}
