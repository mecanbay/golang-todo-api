package repositories

import (
	"golang-todo-api/dto"
	"golang-todo-api/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

type ITodoRepository interface {
	Create(todo *dto.CreateTodoRequest) error
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}
