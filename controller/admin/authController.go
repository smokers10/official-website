package admin

import "github.com/gofiber/fiber/v2"

type authController struct {
	mainController *MainController
}

func (mc *MainController) AuthController() *authController {
	return &authController{mainController: mc}
}

func (ac *authController) LoginPage(c *fiber.Ctx) error {
	return c.Render("admin/login", nil)
}
