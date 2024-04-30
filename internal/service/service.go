package service

import (
	"to-do-checklist/internal/domain"
	"to-do-checklist/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	generatePasswordHash(password string) string
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization Authorization
	TodoList      TodoList
	TodoItem      TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repo.Authorization),
	}
}
