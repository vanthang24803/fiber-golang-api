package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/fiber-api/api/modules/auth"
)

func AuthRouter(app fiber.Router) {
	app.Post("/register", auth.Register)
	app.Post("/login", auth.Login)
}
