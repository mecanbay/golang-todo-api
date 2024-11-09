package handlers

import (
	"golang-todo-api/dto"
	"golang-todo-api/services"
	"golang-todo-api/validators"
	"strconv"

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
		response := &dto.CreateTodoResponse{Status: false, Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err := validators.CreateTodoRequestValidator(request)
	if err != nil {
		response := &dto.CreateTodoResponse{Status: false, Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := h.TodoService.CreateTodo(request)
	if !response.Status {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}

func (h TodoHandler) GetTodoById(c *fiber.Ctx) error {
	Id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := &dto.GetTodoByIdResponse{Status: false, Message: validators.ErrInvalidId.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := h.TodoService.GetTodoById(uint(Id))
	if !response.Status {
		return c.Status(fiber.StatusNotFound).JSON(response)
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	response := h.TodoService.GetAllTodos()
	if !response.Status {
		return c.Status(fiber.StatusNotFound).JSON(response)
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	// Todo Id Params Conversion
	Id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := &dto.UpdateTodoResponse{Status: false, Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Todo Body Parse
	var request dto.UpdateTodoRequest
	if err := c.BodyParser(&request); err != nil {
		response := &dto.UpdateTodoResponse{Status: false, Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Todo Update request validation
	err = validators.UpdateTodoRequestValitator(&request)
	if err != nil {
		response := &dto.UpdateTodoResponse{Status: false, Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := h.TodoService.UpdateTodo(uint(Id), &request)
	if !response.Status {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
