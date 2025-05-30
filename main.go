package main

import (
	"github/ertush/gorest/database"
	"github/ertush/gorest/middleware"
	"github/ertush/gorest/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func loadrouter(app *fiber.App) {

	// Auth

	app.Post("/api/auth/login", router.Login)
	app.Post("/api/auth/register", router.CreateUser)

	// User Routes

	// app.Post("/api/user", middleware.Protected(), router.CreateUser)
	// app.Get("/api/user", middleware.Protected(), router.GetUser)
	app.Get("/api/user/:id", router.GetUser)
	app.Put("/api/user/:id", middleware.Protected(), router.UpdateUser)
	app.Delete("/api/users/:id", middleware.Protected(), router.DeleteUser)

	// Product Routes

	app.Post("/api/product", middleware.Protected(), router.CreateProduct)
	app.Get("/api/products", middleware.Protected(), router.GetProducts)
	app.Get("/api/products/:id", middleware.Protected(), router.GetProduct)
	app.Put("/api/product", middleware.Protected(), router.UpdateProduct)
	app.Delete("/api/products/:id", middleware.Protected(), router.DeleteProduct)

}

func main() {

	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New())
	// middleware.UseAuth(app)

	loadrouter(app)

	log.Fatalln(app.Listen(":4000"))
}
