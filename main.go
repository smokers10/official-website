package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/smokers10/official-website.git/config"
	"github.com/smokers10/official-website.git/router"
)

func main() {
	// get config
	config := config.ReadConfig()

	// init view engine
	viewEngine := html.New("./views", ".html")

	// set fiber config
	fiberConfig := fiber.Config{
		Views: viewEngine,
	}

	// initialize fiber
	app := fiber.New(fiberConfig)

	// call router
	router.RouterGet(app)

	// serve
	app.Listen(config.Application.Port)
}
