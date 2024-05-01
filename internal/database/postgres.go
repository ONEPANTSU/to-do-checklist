package database

import (
	"github.com/jmoiron/sqlx"
	"to-do-checklist/internal/config"
	"to-do-checklist/internal/domain"
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

func (postgres *PostgresDB) CreateAndReturnIDQuery(query string, args ...any) (int, error) {
	var id int
	row := postgres.db.QueryRow(query, args...)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (postgres *PostgresDB) GetOneQuery(model domain.Model, query string, args ...any) ([]interface{}, error) {
	row := postgres.db.QueryRow(query, args...)
	fields := model.GetFields()
	if err := row.Scan(fields...); err != nil {
		return nil, err
	}
	return fields, nil
}
