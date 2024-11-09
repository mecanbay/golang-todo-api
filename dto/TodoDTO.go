package dto

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Desc        string `json:"desc"`
	IsCompleted bool   `json:"is_completed"`
}

type CreateTodoResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"msg"`
}
