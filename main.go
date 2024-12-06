package main

import (
	"log"
	_ "github.com/Hexagonz/back-end-go/database/connection"
	"github.com/Hexagonz/back-end-go/routes"
	"github.com/Hexagonz/back-end-go/utils"
	"github.com/kataras/iris/v12"
)

func main() {
	app, err := routes.RoutePath(iris.New())
	app.Listen(":8080")
    defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
        }
	}()	
	utils.Catch(err)
}

