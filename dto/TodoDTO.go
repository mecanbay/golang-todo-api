package dto

import "golang-todo-api/models"

type CreateTodoRequest struct {
	Title       string `json:"title,omitempty" validate:"required,min=3,max=100"`
	Desc        string `json:"desc"`
	IsCompleted bool   `json:"is_completed"`
}

type CreateTodoResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"msg"`
}

type GetTodoByIdResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"msg"`
	Result  *models.Todo `json:"result"`
}

type GetAllTodosResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"msg"`
	Result  []models.Todo `json:"result"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title,omitempty" validate:"required,min=3,max=100"`
	Desc        string `json:"desc" validate:"required"`
	IsCompleted bool   `json:"is_completed"`
}

type UpdateTodoResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"msg"`
}
