package repository

import (
	"to-do-checklist/internal/database"
	"to-do-checklist/internal/domain"
)

const (
	usersStorage     = "users"
	todoListsStorage = "todo_lists"
	todoItemsStorage = "todo_items"
	userListStorage  = "user_list"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authorization Authorization
	TodoList      TodoList
	TodoItem      TodoItem
}

func NewRepository(db database.Database) *Repository {
	return &Repository{
		Authorization: newAuthRepository(db),
	}
}
