package main

import (
	"github/ertush/gorest/database"
	"github/ertush/gorest/middleware"
	"github/ertush/gorest/views"
	"log"

	"github.com/gofiber/fiber/v2"
)

func loadViews(app *fiber.App) {

	// User Routes

	app.Post("/api/user", views.CreateUser)
	app.Get("/api/users", views.GetUsers)
	app.Get("/api/users/:id", views.GetUser)
	app.Put("/api/user", views.UpdateUser)
	app.Delete("/api/users/:id", views.DeleteUser)

	// Product Routes

	app.Post("/api/product", views.CreateProduct)
	app.Get("/api/products", views.GetProducts)
	app.Get("/api/products/:id", views.GetProduct)
	app.Put("/api/product", views.UpdateProduct)
	app.Delete("/api/products/:id", views.DeleteProduct)

}

func main() {

	database.ConnectDB()

	app := fiber.New()

	middleware.UseAuth(app)

	loadViews(app)

	log.Fatalln(app.Listen(":4000"))
}
