package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/smokers10/official-website.git/controller"
)

func RouterGet(app *fiber.App, session *session.Store) {
	// call controller
	controller := controller.MainController{}

	// index page
	app.Get("/", controller.IndexController().Index)

	// Admin
	adminGroup := app.Group("/admin")

	// Admin - Authentication
	adminGroup.Get("/login", controller.AuthController().LoginPage)

	// Admin - Dashboard
	adminGroup.Get("/home", controller.AdminHomeController().HomePage)
}
