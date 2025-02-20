package main

import (
	"github.com/IsaqueAmorim/rinha-2023/internal/infra/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/pessoas", handlers.CreatePerson)
	app.Get("/pessoas/:id", handlers.FindPersonById)
	app.Get("/pessoas", handlers.FindByText)
	app.Get("/contagem-pessoas", handlers.CountPersons)
	app.Listen(":8080")

}
