package routes

import (
	"golang-todo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App, handler *handlers.TodoHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	todo := v1.Group("/todo")
	todo.Get("/", handler.GetAllTodos)
	todo.Get("/:id", handler.GetTodoById)
	todo.Post("/", handler.CreateTodo)
	todo.Put("/:id", handler.UpdateTodo)
	todo.Delete("/:id", handler.DeleteTodo)
}
