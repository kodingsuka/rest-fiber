package handlers

import (
	"rest-fiber/data"
	"rest-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateTodoHandler(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var input models.UpdateTodoRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}
	var idx int
	for i, todo := range data.Todos {
		if todo.ID == uint(id) {
			idx = i
			break
		}
	}
	data.Todos[idx].Title = input.Title
	data.Todos[idx].IsCompleted = input.IsCompleted
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": nil,
		"data":  data.Todos[idx],
	})
}
