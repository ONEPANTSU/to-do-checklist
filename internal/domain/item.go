package domain

import "errors"

type TodoItem struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Completed   bool   `json:"completed" db:"completed"`
	ListID      int    `json:"list_id" db:"list_id" binding:"required"`
}

type UpdateTodoItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

func (item *UpdateTodoItem) Validate() error {
	if item.Title == nil && item.Description == nil && item.Completed == nil {
		return errors.New("must provide either a title, a description or a completed")
	}
	return nil
}
