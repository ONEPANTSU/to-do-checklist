package service

import (
	"to-do-checklist/internal/config"
	"to-do-checklist/internal/domain"
	"to-do-checklist/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	generatePasswordHash(password string) string
	GenerateToken(authInfo domain.SignIn) (string, error)
	ValidateToken(token string) (int, error)
}

type TodoList interface {
	CreateList(list *domain.TodoList, userID int) (int, error)
	GetAllLists() *[]domain.TodoList
	GetUsersLists(userID int) (*[]domain.TodoList, error)
	GetListById(listID int, userID int) (*domain.TodoList, error)
	UpdateList(list *domain.UpdateTodoList, listID, userID int) error
	DeleteList(listID, userID int) error
}

type TodoItem interface {
	CreateItem(item *domain.TodoItem, userID int) (int, error)
	GetItems(listID, userID int) (*[]domain.TodoItem, error)
	GetItemById(itemID, userID int) (*domain.TodoItem, error)
	UpdateItem(item *domain.UpdateTodoItem, itemID, userID int) error
	DeleteItem(itemID, userID int) error
}

type Service struct {
	Authorization Authorization
	TodoList      TodoList
	TodoItem      TodoItem
}

func NewService(repo *repository.Repository, authConfig *config.AuthConfig) *Service {
	return &Service{
		Authorization: newAuthService(repo.Authorization, authConfig),
		TodoList:      newTodoListService(repo.TodoList),
		TodoItem:      newTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
