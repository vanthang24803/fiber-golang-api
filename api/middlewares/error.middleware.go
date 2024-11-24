package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/fiber-api/pkg/utils"
)

func ErrorHandlingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if err != nil {
			if err.Error() == "not found" {
				return c.Status(fiber.StatusNotFound).JSON(utils.Exception(404, "Not Found!"))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(utils.Exception(500, "Internal Server Error"))
		}

		return nil
	}
}
