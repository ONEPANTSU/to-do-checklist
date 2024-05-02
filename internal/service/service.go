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
}

type TodoItem interface{}

type Service struct {
	Authorization Authorization
	TodoList      TodoList
	TodoItem      TodoItem
}

func NewService(repo *repository.Repository, authConfig *config.AuthConfig) *Service {
	return &Service{
		Authorization: newAuthService(repo.Authorization, authConfig),
		TodoList:      newTodoListService(repo.TodoList),
	}
}
