package services

import (
	"golang-todo-api/dto"
	"golang-todo-api/models"
	"golang-todo-api/repositories"
	"golang-todo-api/validators"
)

type TodoService struct {
	TodoRepository repositories.TodoRepository
}

type ITodoService interface {
	CreateTodo(todo *dto.CreateTodoRequest) *dto.CreateTodoResponse
	GetTodoById(Id uint) *dto.GetTodoByIdResponse
}

func NewTodoService(repository *repositories.TodoRepository) *TodoService {
	return &TodoService{TodoRepository: *repository}
}

func (s TodoService) CreateTodo(request *dto.CreateTodoRequest) *dto.CreateTodoResponse {
	var response dto.CreateTodoResponse

	todo := &models.Todo{
		Title:       request.Title,
		Desc:        request.Desc,
		IsCompleted: request.IsCompleted,
	}

	err := s.TodoRepository.Create(todo)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		return &response
	}

	response.Status = true
	response.Message = "success"
	return &response
}

func (s TodoService) GetTodoById(Id uint) *dto.GetTodoByIdResponse {
	var response dto.GetTodoByIdResponse
	todo, err := s.TodoRepository.GetById(Id)
	if err != nil {
		response.Status = false
		response.Message = validators.ErrTodoNotFound.Error()
		return &response
	}

	response.Status = true
	response.Message = "success"
	response.Result = todo
	return &response
}
