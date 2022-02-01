package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/template/html"
	"github.com/smokers10/official-website.git/libraries"
	"github.com/smokers10/official-website.git/libraries/session"
	"github.com/smokers10/official-website.git/router"
)

func main() {
	// get config
	appConfig := libraries.ReadConfig().Application

	// init view engine
	viewEngine := html.New("./views", ".html")

	// set fiber config
	fiberConfig := fiber.Config{
		Views: viewEngine,
	}

	// initialize fiber
	app := fiber.New(fiberConfig)

	// initialize static folder
	app.Static("/", "./public")

	// initialize fiber neccessary middleware
	app.Use(csrf.New())
	app.Use(recover.New())

	// call router
	router.RouterGet(app, session.Store().MYSQL())

	// serve
	app.Listen(appConfig.Port)
}
