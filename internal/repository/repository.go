package repository

import (
	"to-do-checklist/internal/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username string) (*domain.User, error)
}

type TodoList interface {
	CreateList(list *domain.TodoList, userID int) (int, error)
	GetAllLists() *[]domain.TodoList
	GetUsersLists(userID int) (*[]domain.TodoList, error)
	GetListByID(listID int, userID int) (*domain.TodoList, error)
	UpdateList(list *domain.UpdateTodoList, listID, userID int) error
	DeleteList(userID int, listID int) error
}

type TodoItem interface{}

type Repository struct {
	Authorization Authorization
	TodoList      TodoList
	TodoItem      TodoItem
}
