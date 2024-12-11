package controllers

import (
	"github.com/Hexagonz/back-end-go/middleware/jwttoken"
	"github.com/kataras/iris/v12"
)

func RefreshToken(ctx iris.Context) {
	token, err := jwttoken.RefreshTokenJWT(ctx)
	if err != nil {
		ctx.JSON(&ErrorResponse{
			Status: "error",
			Message: "Failed to generate token",
			Errors: err,
		})
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(&Response{
		Status: "success",
		Data: &RefToken{
			AccessToken: token.AccessToken,
			ExpiredAt: token.ExpiredAt,
		},
		Message: "Generate token successfully...",
	})
}
