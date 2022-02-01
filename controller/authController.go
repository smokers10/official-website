package controller

import "github.com/gofiber/fiber/v2"

type authController struct{}

func (mc *MainController) AuthController() *authController {
	return &authController{}
}

func (ac *authController) LoginPage(c *fiber.Ctx) error {
	return c.Render("admin/login", nil)
}
