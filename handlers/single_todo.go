package handlers

import (
	"rest-fiber/data"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetSingleTodoHandler(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	for _, todo := range data.Todos {
		if todo.ID == uint(id) {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"data": todo,
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "data not found",
	})
}
