package controllers

import (
	"fmt"

	"github.com/Hexagonz/back-end-go/middleware/jwttoken"
	"github.com/kataras/iris/v12"
)

func RefreshToken(ctx iris.Context) {
	token, err := jwttoken.RefreshTokenJWT(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"access_token": token.AccessToken,
		"expired_at":   token.ExpiredAt,
	})
}
