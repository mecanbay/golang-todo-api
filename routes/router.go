package routes

import (
	"golang-todo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	todo := v1.Group("/todo")
	todo.Get("/", handlers.CreateTodo)
}
