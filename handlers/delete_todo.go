package handlers

import (
	"rest-fiber/data"
	"rest-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteTodoHandler(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	for i, todo := range data.Todos {
		if todo.ID == uint(id) {
			// [1, 3, 4, 4]
			//  0  1  2  3
			copy(data.Todos[i:], data.Todos[i+1:])
			data.Todos[len(data.Todos)-1] = models.Todo{}
			data.Todos = data.Todos[:len(data.Todos)-1]
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": "data deleted",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "todo not found",
	})
}
