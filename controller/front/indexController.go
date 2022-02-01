package front

import "github.com/gofiber/fiber/v2"

type indexController struct{}

func (mc *MainController) IndexController() *indexController {
	return &indexController{}
}

func (idx *indexController) Index(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
