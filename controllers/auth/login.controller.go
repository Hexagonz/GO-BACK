package controllers

import (
	"fmt"
	"github.com/Hexagonz/back-end-go/database"
	"github.com/Hexagonz/back-end-go/models"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	validate = validator.New()
	validate.RegisterStructValidation(userLoginValidation, Users{})
}

func Login(ctx iris.Context) {
	db, errs = database.SetupDatabase()
	if errs != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "Database connection error",
			"error":   errs.Error(),
		})
		return
	}

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

	var existingUser Users
	user_check := db.Where("email = ?", user.Email).First(&existingUser)
	if user_check.Error != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "Database error occurred",
			"error":   user_check.Error.Error(),
		})
		return
	}

	if user_check.RowsAffected == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "User not found",
			"error":   "Email or Password is incorrect",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "Invalid password",
			"error":   "Email or Password is incorrect",
		})
		return
	}

	token, err := GenerateJWT(user.Email, user.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "Failed to generate token",
			"errors":   err.Error(),
		})
		return
	}
	
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"status":  "success",
		"data":    existingUser,
		"token": token,
		"message": "User login successfully",
	})
}

func userLoginValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(models.Users)
	if len(user.Email) == 0 {
		sl.ReportError(user.Email, "Email", "email", "required", "")
	}
	if len(user.Password) == 0 {
		sl.ReportError(user.Password, "Password", "password", "required", "")
	}
}
