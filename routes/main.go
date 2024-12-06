package routes

import (
	"github.com/Hexagonz/back-end-go/routes/v1/auth"
	"github.com/kataras/iris/v12"
)

func RoutePath(app *iris.Application) (*iris.Application, error) {
	pathApi := app.Party("/api/v1")
	auth.RoutesAuth(pathApi)
	return app, nil
}