package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	jwt "github.com/smokers10/official-website.git/libraries/jsonwebtoken"
)

func AdminOnly(session *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := session.Get(c)
		if err != nil {
			panic(err)
		}

		// ambil token dari session
		token := sess.Get("token")

		// check token
		if token == nil {
			return c.Redirect("/admin/login")
		}

		// verifikasi token
		claims, err := jwt.Init(&jwt.Payload{}).Verificate(token.(string))
		if err != nil {
			panic(err)
		}

		// assign claims ke c.locals()
		c.Locals("id", claims.ID)
		c.Locals("email", claims.Email)
		c.Locals("is_logged", claims.IsLogged)

		return c.Next()
	}
}

func Guest(session *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := session.Get(c)
		if err != nil {
			panic(err)
		}

		// ambil token dari session
		token := sess.Get("token")

		// redirect ke home jika sudah ada token
		if token != nil {
			return c.Redirect("/admin/home")
		}

		return c.Next()
	}
}
