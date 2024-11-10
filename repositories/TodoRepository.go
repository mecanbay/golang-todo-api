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
	Update(todo *dto.UpdateTodoRequest) error
	Delete(todo *models.Todo) error
	GetById(Id uint) (*models.Todo, error)
	GetAll() ([]models.Todo, error)
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r TodoRepository) GetById(Id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, Id)
	return &todo, err.Error
}

func (r TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r TodoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r TodoRepository) Delete(todo *models.Todo) error {
	return r.db.Delete(todo).Error
}
