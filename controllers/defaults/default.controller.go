package controllers

import (
	"fmt"

	"github.com/Hexagonz/back-end-go/middleware/jwttoken"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func Default(ctx iris.Context) {
	claims := jwt.Get(ctx).(*jwttoken.Claims)
	fmt.Println(claims.ID)
	fmt.Println(claims.Email)
	ctx.JSON(iris.Map{
		"status":  "success",
		"message": "User registered successfully",
	})
}
