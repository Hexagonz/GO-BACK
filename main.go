package main

import (
	"log"

	_ "github.com/Hexagonz/back-end-go/database/connection"
	private_routes "github.com/Hexagonz/back-end-go/routes/v1/private_routes/default_routes"
	"github.com/Hexagonz/back-end-go/middleware"
	"github.com/Hexagonz/back-end-go/routes/v1/auth"
	"github.com/Hexagonz/back-end-go/utils"
	"github.com/kataras/iris/v12"
)

const (
	path string = "/api/v1"
)

func routePath(app *iris.Application) (*iris.Application, error) {
	pathApi := app.Party(path)
	auth.RoutesAuth(pathApi)
	app.Use(middleware.AuthMiddleware)
	privateApi := app.Party(path, middleware.Validate)
	private_routes.RoutesAuth(privateApi)
	return app, nil
}

func main() {
	app, err := routePath(utils.App)
	app.Listen(":8080")
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	utils.Catch(err)
}
