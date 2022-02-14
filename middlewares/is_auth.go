package middlewares

import (
	"rest-fiber/data"
	"rest-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func IsAuth(c *fiber.Ctx) error {
	token := c.Cookies("token", "")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token not provided",
		})
	}
	user := models.User{}
	for _, u := range data.Users {
		if u.Token == token {
			user = u
		}
	}
	if user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	c.Locals("userid", user.ID)
	return c.Next()
}
