package models

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
	UserID      uint   `json:"-"`
}
