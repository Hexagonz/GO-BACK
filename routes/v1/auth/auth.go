package auth

import (
	"github.com/Hexagonz/back-end-go/controllers/auth"
	"github.com/kataras/iris/v12/core/router"
)

func RoutesAuth(app router.Party) {
	app.Post("/register", controllers.Register)
}

