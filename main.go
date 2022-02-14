package main

import (
	"rest-fiber/data"
	"rest-fiber/handlers"
	"rest-fiber/middlewares"
	"rest-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	data.Users = append(data.Users, models.User{
		ID:    1,
		Name:  "Andika",
		Token: "abc",
	}, models.User{
		ID:    2,
		Name:  "John",
		Token: "def",
	})
	data.Todos = append(data.Todos, models.Todo{
		ID:          uint(data.LastTodoID),
		Title:       "Belajar Golang",
		IsCompleted: false,
		UserID:      1,
	})
	data.LastTodoID++

	app := fiber.New()
	app.Use(middlewares.IsAuth)

	app.Get("/", handlers.GetAllTodosHandler)

	app.Post("/", handlers.CreateTodoHandler)

	app.Get("/:id", middlewares.IsOwner, handlers.GetSingleTodoHandler)

	app.Put("/:id", middlewares.IsOwner, handlers.UpdateTodoHandler)

	app.Delete("/:id", middlewares.IsOwner, handlers.DeleteTodoHandler)

	app.Listen(":8080")
}
