package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/smokers10/official-website.git/controller/admin"
	"github.com/smokers10/official-website.git/controller/front"
)

func RouterGet(app *fiber.App, session *session.Store) {
	// call controller
	controllerAdmin := admin.MainController{}
	controllerFront := front.MainController{}

	// index page
	app.Get("/", controllerFront.IndexController().Index)

	// Admin
	adminGroup := app.Group("/admin")

	// Admin - Authentication
	adminGroup.Get("/login", controllerAdmin.AuthController().LoginPage)

	// Admin - Dashboard
	adminGroup.Get("/home", controllerAdmin.AdminHomeController().HomePage)
}
