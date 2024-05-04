package postgres

import (
	"github.com/jmoiron/sqlx"
	"to-do-checklist/internal/repository"
)

func NewPostgresRepository(db *sqlx.DB) *repository.Repository {
	return &repository.Repository{
		Authorization: newAuthRepository(db),
		TodoList:      newTodoListsRepository(db),
		TodoItem:      newTodoItemRepository(db),
	}
}
