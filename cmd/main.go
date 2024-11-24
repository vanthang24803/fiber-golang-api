package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/fiber-api/api/middlewares"
	"github.com/vanthang24803/fiber-api/api/router"
	"github.com/vanthang24803/fiber-api/internal/config"
	"github.com/vanthang24803/fiber-api/internal/database"
)

func main() {

	config.LoadEnvFile()
	database.ConnectionDB()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "3000"
	}

	app := fiber.New()

	app.Use(middlewares.LoggerMiddleware)
	app.Use(middlewares.ErrorHandlingMiddleware())

	api := app.Group("/api")

	router.AuthRouter(api)
	router.NotFoundRoute(app)

	app.Listen(":" + PORT)

	if err := app.Listen(":" + PORT); err != nil {
		log.Fatal("Error starting server:", err)
	}

	log.Printf("Application running on port %v\n ✅", PORT)
}
