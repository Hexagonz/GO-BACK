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

const path string = "/api/v1"

func routePath(app *iris.Application) (*iris.Application, error) {
	pathApi := app.Party(path)
	auth.RoutesAuth(pathApi)
	//middleware
	app.Use(middleware.AuthMiddleware)
	// protected api
	protectedApi := app.Party(path, middleware.Validate)
	private_routes.RoutesAuth(protectedApi)
	return app, nil
}

func main() {
	app := utils.App
	routePath(utils.App)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
}
