package controller

import "github.com/gofiber/fiber/v2"

type adminHomeController struct{}

func (mc *MainController) AdminHomeController() *adminHomeController {
	return &adminHomeController{}
}

func (ahc *adminHomeController) HomePage(c *fiber.Ctx) error {
	return c.Render("admin/home", nil)
}
