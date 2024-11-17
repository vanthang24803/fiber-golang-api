package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vanthang24803/fiber-api/internal/database"
)

func FindOne(c *fiber.Ctx) error {
	db := database.ConnectionDB()
	defer db.Close()

	userIDParam := c.Params("userID")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := findOne(db, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func FindAll(c *fiber.Ctx) error {
	db := database.ConnectionDB()
	defer db.Close()

	users, err := findAll(db)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}
