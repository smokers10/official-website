package controller

import (
	"github.com/gofiber/fiber/v2"
)

type adminHomeController struct {
	MainController *MainController
}

func (mc *MainController) AdminHomeController() *adminHomeController {
	return &adminHomeController{MainController: mc}
}

func (a *adminHomeController) HomePage(c *fiber.Ctx) error {
	return c.Render("admin/home", nil, "layouts/admin")
}
