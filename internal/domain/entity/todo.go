package entity

import "time"

type Todo struct {
	ID        int
	Title    string
	Content     string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type ListTodos []Todo

func NewRegisterTodo(
	title string,
	content string,
	createdBy string,
	updatedBy string,
) *Todo {
	return &Todo{
		Title:    title,
		Content:      content,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}
}
