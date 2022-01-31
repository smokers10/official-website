package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/official-website.git/controller"
)

func RouterGet(app *fiber.App) {
	// call controller
	controller := controller.MainController{}

	// index page
	app.Get("/", controller.IndexController().Index)
}
