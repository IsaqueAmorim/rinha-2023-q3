package handlers

import (
	"github.com/IsaqueAmorim/rinha-2023/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

func FindPersonById(c *fiber.Ctx) error {
	id := c.Params("id")

	person, err := service.GetPersonById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if person == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Person not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(person)
}

func FindByText(c *fiber.Ctx) error {
	t := c.Query("t")

	if t == "" {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "missing query param 't'"})
	}

	persons, err := service.FindByText(t)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&persons)
}
