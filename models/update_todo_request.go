package models

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}
