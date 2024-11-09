package validators

import "errors"

var ErrTodoTitle = errors.New("title is required and must be between 3 and 100 characters")
var ErrTodoDesc = errors.New("desc is required")
var ErrTodoIsCompleted = errors.New("is_completed is required")

var ErrInvalidId = errors.New("invalid id")
var ErrTodoNotFound = errors.New("todo not found")
