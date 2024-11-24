package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/fiber-api/pkg/utils"
)

func NotFoundRoute(app *fiber.App) {
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(utils.Exception(404, "Not Found!"))
	})
}
