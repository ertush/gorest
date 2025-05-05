package main

import (
	"github/ertush/gorest/database"
	"github/ertush/gorest/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.Status(200).SendString("Hello from go fiber")
}

func loadRoutes(app *fiber.App) {

	app.Get("/", welcome)
	app.Post("/api/user", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/user", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

}

func main() {

	database.ConnectDB()

	app := fiber.New()

	loadRoutes(app)

	log.Fatalln(app.Listen(":4000"))
}
