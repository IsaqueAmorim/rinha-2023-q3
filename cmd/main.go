package main

import (
	"github.com/IsaqueAmorim/rinha-2023/internal/infra/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/person", handlers.CreatePerson)
	app.Listen(":8080")

}
