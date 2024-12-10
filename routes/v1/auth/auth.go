package auth

import (
	controllers "github.com/Hexagonz/back-end-go/controllers/auth"
	"github.com/Hexagonz/back-end-go/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func RoutesAuth(app router.Party) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Use(middleware.AuthMiddleware)
	app.Use(middleware.ValidateRefresh)
	app.Post("/refresh", controllers.RefreshToken)
}
