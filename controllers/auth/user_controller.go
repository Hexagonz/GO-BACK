package controllers

import (
	"fmt"

	"github.com/Hexagonz/back-end-go/database"
	"github.com/Hexagonz/back-end-go/models"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterStructValidation(UserStructLevelValidation, models.Users{})
}

func Register(ctx iris.Context) {
	var user models.Users
	db, errs := database.SetupDatabase()

	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err = validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{
				"status":  "error",
				"message": "Invalid validation error",
				"error":   err.Error(),
			})
			return
		}

		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Field '%s' failed on the '%s' rule", err.Field(), err.Tag())
		}

		ctx.StatusCode(422)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "Validation failed",
			"errors":  errors,
		})
		return
	}
	user_check := db.Where("name = ? ", user.Name).First(&user)

	if user_check.RowsAffected > 0 {
		ctx.StatusCode(iris.StatusConflict)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "User with this name already exists",
		})
		return
	}

	email_check := db.Where("email = ? ", user.Email).First(&user)

	if email_check.RowsAffected > 0 {
		ctx.StatusCode(iris.StatusConflict)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "User with this email already exists",
		})
		return
	}
	user_create := models.Users{Name: user.Name, Email: user.Email, Password: user.Password}
	db.Create(&user_create)

	if errs != nil {
		panic(fmt.Sprintf("Errors, %v", errs))
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"status":  "success",
		"data":    user_create,
		"message": "User registered successfully",
	})
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(models.Users)
	if len(user.Name) == 0 {
		sl.ReportError(user.Name, "Name", "username", "required", "")
	}
	if len(user.Email) == 0 {
		sl.ReportError(user.Email, "Email", "email", "required", "")
	}
	if len(user.Password) == 0 {
		sl.ReportError(user.Password, "Password", "password", "required", "")
	}
}
