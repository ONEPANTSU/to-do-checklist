package database

import (
	"github.com/jmoiron/sqlx"
	"to-do-checklist/internal/config"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (postgres *PostgresDB) Connect(cfg *config.DBConfig) error {
	db, err := sqlx.Open(cfg.DBDriver, cfg.GetConnectionURL())
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	postgres.db = db
	return nil
}

func (postgres *PostgresDB) CreateQuery(query string, args ...any) (int, error) {
	var id int
	row := postgres.db.QueryRow(query, args...)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
