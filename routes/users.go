package routes

import (
	"github/ertush/gorest/database"
	"github/ertush/gorest/models"
	"github/ertush/gorest/serializers"

	"github.com/gofiber/fiber/v2"
)

type ResponseError struct {
	Message string `json:"message"`
}

func CreateResponseUser(user models.User) serializers.User {
	return serializers.User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(201).JSON(responseUser)

}

func GetUser(c *fiber.Ctx) error {

	var user models.User

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 && user.FirstName == "" {
		return c.Status(400).JSON(ResponseError{Message: "User not found"})
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {

	var users []models.User
	var responseUsers []serializers.User

	database.Database.Db.Find(&users)

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)

}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.UpdateColumns(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Delete(&user, "id = ?", id)

	return c.Status(200).JSON(ResponseError{Message: "User deleted successfully"})
}
