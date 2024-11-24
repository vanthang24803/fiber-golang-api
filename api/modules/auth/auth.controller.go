package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/fiber-api/api/modules/auth/common"
	"github.com/vanthang24803/fiber-api/internal/database"
)

func Register(c *fiber.Ctx) error {

	db := database.ConnectionDB()
	defer db.Close()

	var req common.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := register(db, &req)
	return c.Status(201).JSON(result)

}

func Login(c *fiber.Ctx) error {
	db := database.ConnectionDB()
	defer db.Close()

	var req common.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := login(db, &req)
	return c.Status(201).JSON(result)
}
