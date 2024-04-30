package repository

import (
	"fmt"
	"to-do-checklist/internal/database"
	"to-do-checklist/internal/domain"
)

type AuthRepository struct {
	db database.Database
}

func newAuthRepository(db database.Database) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user domain.User) (int, error) {
	query := fmt.Sprintf(
		"insert into %s (username, email, hashed_password) "+
			"values ($1, $2, $3) returning id",
		usersStorage,
	)
	return r.db.CreateQuery(query, user.Username, user.Email, user.Password)
}
