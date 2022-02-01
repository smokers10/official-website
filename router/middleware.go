package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
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

		return c.Next()
	}
}
