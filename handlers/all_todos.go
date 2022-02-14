package handlers

import (
	"rest-fiber/data"
	"rest-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllTodosHandler(c *fiber.Ctx) error {
	todos := []models.Todo{}
	userid := c.Locals("userid").(uint)
	for _, t := range data.Todos {
		if t.UserID == userid {
			todos = append(todos, t)
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todos,
	})
}
