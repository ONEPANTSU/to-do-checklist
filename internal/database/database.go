package database

import (
	"to-do-checklist/internal/config"
	"to-do-checklist/internal/domain"
)

type Database interface {
	Connect(cfg *config.DBConfig) error
	CreateAndReturnIDQuery(query string, args ...any) (int, error)
	GetOneQuery(model domain.Model, query string, args ...any) ([]interface{}, error)
}
