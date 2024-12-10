package middleware

import (
	"strings"

	"github.com/Hexagonz/back-end-go/middleware/jwttoken"
	"github.com/kataras/iris/v12"
)

var (
	Validate = jwttoken.Verifier.Verify(func() interface{} {
		return new(jwttoken.Claims)
	})
	ValidateRefresh = jwttoken.VerifierRefresh.Verify(func() interface{} {
		return new(jwttoken.Claims)
	})
)

func AuthMiddleware(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"errors": "Authorization header is missing"})
		return
	}

	if parts := strings.Split(token, " "); len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"errors": "Invalid authorization header format"})
		return
	}
	ctx.Next()
}
