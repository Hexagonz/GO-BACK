package controllers

import (
	"fmt"

	"github.com/Hexagonz/back-end-go/database"
	"github.com/Hexagonz/back-end-go/models"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

func init() {
	validate = validator.New()
	validate.RegisterStructValidation(userRegisterValidation, RegisterUser{})
}

func Register(ctx iris.Context) {
	db, errs = database.SetupDatabase()
	err := ctx.ReadJSON(&register)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
			Errors:  err.Error(),
		})
		return
	}

	err = validate.Struct(register)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(&ErrorResponse{
				Status:  "error",
				Message: "Invalid validation error",
				Errors:  err.Error(),
			})
			return
		}

		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "eqfield" && err.Field() == "Password_Confirmation" {
				errors[err.Field()] = fmt.Sprintf("Field '%s' failed on the '%s' rule password not mismatch password_confrimation", err.Field(), err.Tag())
			}
			errors[err.Field()] = fmt.Sprintf("Field '%s' failed on the '%s' rule", err.Field(), err.Tag())
		}

		ctx.StatusCode(422)
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Validation failed",
			Errors:  errors,
		})
		return
	}
	user_check := db.Where("name = ? ", register.Name).First(&user).Error

	if user_check == nil {
		ctx.StatusCode(iris.StatusConflict)
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Validate failed",
			Errors:  map[string]interface{}{"email": "Email already exists"},
		})
		return
	}

	email_check := db.Where("email = ? ", register.Email).First(&user).Error
	if email_check == nil {
		ctx.StatusCode(iris.StatusConflict)
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Validate failed",
			Errors:  map[string]interface{}{"email": "Email already exists"},
		})
		return
	}
	user_create := models.Users{Name: register.Name, Email: register.Email, Password: register.Password}
	db.Create(&user_create)

	if errs != nil {
		panic(fmt.Sprintf("Errors, %v", errs))
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(&Response{
		Status:  "success",
		Data:    user_create,
		Message: "User registered successfully",
	})
}

func userRegisterValidation(sl validator.StructLevel) {
	register := sl.Current().Interface().(RegisterUser)
	if len(register.Name) == 0 {
		sl.ReportError(register.Name, "Name", "name", "required", "")
	}
	if len(register.Email) == 0 {
		sl.ReportError(register.Email, "Email", "email", "required", "")
	}
	if len(register.Password) == 0 {
		sl.ReportError(register.Password, "Password", "password", "required", "")
	}
	if len(register.Password_Confirmation) == 0 {
		sl.ReportError(register.Password_Confirmation, "Password", "password_confirmation", "required", "")
	}
}
