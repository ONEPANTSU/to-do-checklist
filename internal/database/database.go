package database

import "to-do-checklist/internal/config"

type Database interface {
	Connect(cfg *config.DBConfig) error
	CreateQuery(query string, args ...any) (int, error)
}
