package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Hexagonz/back-end-go/middleware/jwttoken"
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
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
			Errors:  err.Error(),
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
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Validation failed",
			Errors:  errors,
		})
		return
	}

	var existingUser models.Users
	user_check := db.Where("email = ?", user.Email).First(&existingUser)
	if user_check.Error != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Database error occurred",
			Errors:  user_check.Error.Error(),
		})
		return
	}

	if user_check.RowsAffected == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "User not found",
			Errors:  "Email or Password is incorrect",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Invalid password",
			Errors:  "Email or Password is incorrect",
		})
		return
	}
	str := strconv.FormatUint(uint64(existingUser.ID), 10)
	accses_token, err := jwttoken.GenerateTokenJwt(user.Email, str, ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(&ErrorResponse{
			Status:  "error",
			Message: "Failed to generate token",
			Errors:  err.Error(),
		})
		return
	}

	sessionUser := models.RefreshToken{
		Refresh_Token: accses_token.AccessToken,
		UserAgent:     ctx.GetHeader("User-Agent"),
		ExpiredAt:     time.Now().Add(time.Duration(accses_token.RefreshExpiresAt)),
	}
	db.Create(&sessionUser)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(&Response{
		Status:  "success",
		Data:    accses_token,
		Message: "User login successfully",
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
