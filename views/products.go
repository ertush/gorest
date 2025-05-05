package views

import (
	"github/ertush/gorest/database"
	"github/ertush/gorest/models"
	"github/ertush/gorest/serializers"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Name         string         `json:"name"`
	SerialNumber string         `json:"serial_number"`
	Product      models.Product `json:"product"`
}

func CreateResponseProduct(product models.Product) serializers.Product {
	return serializers.Product{ID: product.ID, Name: product.Name, SerialNumber: product.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {

	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(201).JSON(responseProduct)

}

func GetProduct(c *fiber.Ctx) error {

	var product models.Product

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0 && product.Name == "" {
		return c.Status(400).JSON(ResponseError{Message: "Product not found"})
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {

	var products []models.Product
	var responseProducts []serializers.Product

	database.Database.Db.Find(&products)

	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)

}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.UpdateColumns(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)

}

func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Delete(&product, "id = ?", id)

	return c.Status(200).JSON(ResponseError{Message: "Product deleted successfully"})
}
