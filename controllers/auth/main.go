package controllers

import (
	"github.com/Hexagonz/back-end-go/database"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var (
	db, errs = database.SetupDatabase()
)

var user Users
var register RegisterUser

type RegisterUser struct {
	Name                  string `json:"name" validate:"required,min=5,max=30"`
	Email                 string `json:"email" validate:"required,email"`
	Password              string `json:"password" validate:"required,min=8,max=30"`
	Password_Confirmation string `json:"password_confirmation" validate:"required,min=8,max=30,eqfield=Password"`
}

type Users struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

type JWTClaim struct {
	Name  string `json:"id"`
	Email string `json:"email"`
}

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type RefToken struct {
	AccessToken string `json:"access_token"`
	ExpiredAt int64 `json:"expired_at"`
}