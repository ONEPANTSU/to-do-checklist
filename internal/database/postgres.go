package database

import (
	"github.com/jmoiron/sqlx"
	"to-do-checklist/internal/config"
)

func Connect(cfg *config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.DBDriver, cfg.GetConnectionURL())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
