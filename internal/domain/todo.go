package domain

import "errors"

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	ListID int `json:"list_id"`
}

type UpdateTodoList struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (list *UpdateTodoList) Validate() error {
	if list.Title == nil && list.Description == nil {
		return errors.New("must provide either a title, or a description")
	}
	return nil
}

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	ListID      int    `json:"list_id"`
}
