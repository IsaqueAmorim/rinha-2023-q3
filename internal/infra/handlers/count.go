package handlers

import (
	"fmt"
	"os"

	"github.com/IsaqueAmorim/rinha-2023/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

func CountPersons(c *fiber.Ctx) error {
	_, err := service.CountPersons()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println(os.Getpid())

	host, _ := os.Hostname()

	return c.Status(fiber.StatusOK).JSON(host)
}
