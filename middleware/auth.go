package middleware

import (
	"github/ertush/gorest/views"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func UseAuth(app *fiber.App) {

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("USER"): os.Getenv("PASS"),
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(views.ResponseError{Message: "Unauthorized User"})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))
}
