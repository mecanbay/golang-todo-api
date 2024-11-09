package validators

import (
	"golang-todo-api/dto"

	"github.com/go-playground/validator/v10"
)

func CreateTodoRequestValidator(request *dto.CreateTodoRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				return ErrTodoTitle
			case "Desc":
				return ErrTodoDesc
			}

		}
	}
	return nil
}

func UpdateTodoRequestValitator(request *dto.UpdateTodoRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				return ErrTodoTitle
			case "Desc":
				return ErrTodoDesc
			case "IsCompleted":
				return ErrTodoIsCompleted
			}
		}
	}
	return nil
}
