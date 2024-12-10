package private_routes

import (
	"github.com/Hexagonz/back-end-go/controllers/defaults"

	"github.com/kataras/iris/v12/core/router"
)

func RoutesAuth(app router.Party) {
	app.Get("/",controllers.Default)
}
