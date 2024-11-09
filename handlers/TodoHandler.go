package handlers

import (
	"golang-todo-api/dto"
	"golang-todo-api/services"
	"golang-todo-api/validators"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	TodoService services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: *service}
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var request *dto.CreateTodoRequest
	if err := c.BodyParser(&request); err != nil {
		response := &dto.CreateTodoResponse{
			Status:  false,
			Message: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err := validators.CreateTodoRequestValidator(request)
	if err != nil {
		response := &dto.CreateTodoResponse{
			Status:  false,
			Message: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := h.TodoService.CreateTodo(request)
	if !response.Status {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}
