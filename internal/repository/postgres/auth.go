package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"to-do-checklist/internal/domain"
)

type AuthRepository struct {
	db *sqlx.DB
}

func newAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user domain.User) (int, error) {
	query := fmt.Sprintf(
		"insert into %s (username, email, hashed_password) "+
			"values ($1, $2, $3) returning id",
		usersTable,
	)
	var id int
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepository) GetUser(username string) (*domain.User, error) {
	query := fmt.Sprintf(
		"select * from %s where username = $1",
		usersTable,
	)
	user := domain.User{}
	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
