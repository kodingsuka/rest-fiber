package handlers

import (
	"rest-fiber/data"
	"rest-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTodoHandler(c *fiber.Ctx) error {
	userid := c.Locals("userid").(uint)
	var input models.CreateTodoRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}
	data.Todos = append(data.Todos, models.Todo{
		ID:          uint(data.LastTodoID),
		Title:       input.Title,
		IsCompleted: false,
		UserID:      userid,
	})
	data.LastTodoID++
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   nil,
		"message": "todo created",
	})
}
