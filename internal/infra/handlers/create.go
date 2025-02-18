package handlers

import (
	"github.com/IsaqueAmorim/rinha-2023/internal/domain/entity"
	"github.com/IsaqueAmorim/rinha-2023/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

func CreatePerson(c *fiber.Ctx) error {
	person, err := parseJsonToPerson(c)

	if err != nil {
		return err
	}

	exists, err := service.CheckIfNicknameExists(person.Nickname)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if exists {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "Nickname already exists",
		})
	}

	id, err := service.CreatePerson(person)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Response().Header.Add("Location", c.BaseURL()+"/person/"+id)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":      id,
		"message": "Person created successfully",
	})
}

func parseJsonToPerson(c *fiber.Ctx) (*entity.Person, error) {
	var person *entity.Person

	if err := c.BodyParser(&person); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return person, nil
}

func checkIfNicknameExists(c *fiber.Ctx, nickname string) error {

	return nil
}
