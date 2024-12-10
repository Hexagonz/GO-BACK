package controllers

import (
	"github.com/go-playground/validator/v10"
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
	ID 		 string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

type JWTClaim struct {
	Name  string `json:"id"`
	Email string `json:"email"`
}
