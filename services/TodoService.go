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
	UpdateTodo(todo *dto.UpdateTodoRequest) *dto.UpdateTodoResponse
	DeleteTodo(Id uint) *dto.DeleteTodoResponse
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

func (s TodoService) UpdateTodo(Id uint, request *dto.UpdateTodoRequest) *dto.UpdateTodoResponse {
	todo, err := s.TodoRepository.GetById(Id)
	if err != nil {
		return &dto.UpdateTodoResponse{Status: false, Message: err.Error()}
	}

	todo.Title = request.Title
	todo.Desc = request.Desc
	todo.IsCompleted = request.IsCompleted

	err = s.TodoRepository.Update(todo)
	if err != nil {
		return &dto.UpdateTodoResponse{Status: false, Message: err.Error()}
	}
	return &dto.UpdateTodoResponse{Status: true, Message: "success"}
}

func (s TodoService) DeleteTodo(Id uint) *dto.DeleteTodoResponse {
	todo, err := s.TodoRepository.GetById(uint(Id))
	if err != nil {
		return &dto.DeleteTodoResponse{Status: false, Message: err.Error()}
	}

	err = s.TodoRepository.Delete(todo)
	if err != nil {
		return &dto.DeleteTodoResponse{Status: false, Message: err.Error()}
	}
	return &dto.DeleteTodoResponse{Status: true, Message: "success"}
}
