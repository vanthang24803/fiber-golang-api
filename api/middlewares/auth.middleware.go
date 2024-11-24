package middlewares

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vanthang24803/fiber-api/pkg/utils"
)

func AuthenticationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")

		if token == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.Exception(401, "Unauthorized"))
		}

		if len(token) > 6 && token[:7] == "Bearer " {
			token = token[7:]
		}

		secretKey := os.Getenv("JWT_SECRET")
		_, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			log.Printf("Error parsing token: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(utils.Exception(401, "Unauthorized"))
		}

		return c.Next()
	}
}

func AuthorizationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
