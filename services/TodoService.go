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
	GetAllTodos() *dto.GetAllTodosResponse
}

func NewTodoService(repository *repositories.TodoRepository) *TodoService {
	return &TodoService{TodoRepository: *repository}
}

func (s TodoService) CreateTodo(request *dto.CreateTodoRequest) *dto.CreateTodoResponse {
	todo := &models.Todo{
		Title:       request.Title,
		Desc:        request.Desc,
		IsCompleted: request.IsCompleted,
	}

	err := s.TodoRepository.Create(todo)
	if err != nil {
		return &dto.CreateTodoResponse{Status: false, Message: err.Error()}
	}

	return &dto.CreateTodoResponse{Status: true, Message: "success"}
}

func (s TodoService) GetTodoById(Id uint) *dto.GetTodoByIdResponse {
	todo, err := s.TodoRepository.GetById(Id)
	if err != nil {
		return &dto.GetTodoByIdResponse{Status: false, Message: validators.ErrTodoNotFound.Error()}
	}
	return &dto.GetTodoByIdResponse{Status: true, Message: "success", Result: todo}
}

func (s TodoService) GetAllTodos() *dto.GetAllTodosResponse {
	todos, err := s.TodoRepository.GetAll()
	if err != nil {
		return &dto.GetAllTodosResponse{Status: false, Message: err.Error()}
	}
	return &dto.GetAllTodosResponse{Status: true, Message: "success", Result: todos}
}
