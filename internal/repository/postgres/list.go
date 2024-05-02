package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"to-do-checklist/internal/domain"
)

type TodoListsRepository struct {
	db *sqlx.DB
}

func newTodoListsRepository(db *sqlx.DB) *TodoListsRepository {
	return &TodoListsRepository{db: db}
}

func (r *TodoListsRepository) CreateList(list *domain.TodoList, userID int) (int, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf(
		"insert into %s (title, description) "+
			"values ($1, $2) returning id",
		todoListsTable,
	)
	row := transaction.QueryRow(query, list.Title, list.Description)
	var listID int
	if err := row.Scan(&listID); err != nil {
		_ = transaction.Rollback()
		return 0, err
	}
	query = fmt.Sprintf(
		"insert into %s (user_id, list_id) "+
			"values ($1, $2) returning id",
		userListTable,
	)
	if _, err = transaction.Exec(query, userID, listID); err != nil {
		_ = transaction.Rollback()
		return 0, err
	}
	return listID, transaction.Commit()
}

func (r *TodoListsRepository) GetAllLists() *[]domain.TodoList {
	query := "select * from todo_lists order by id desc"
	var lists []domain.TodoList
	err := r.db.Select(&lists, query)
	if err != nil {
		return nil
	}
	return &lists
}

func (r *TodoListsRepository) GetUsersLists(userID int) (*[]domain.TodoList, error) {
	query := fmt.Sprintf(
		"select list.id, list.title, list.description "+
			"from %s list inner join %s user_list "+
			"on user_list.list_id = list.id "+
			"where user_list.user_id = $1",
		todoListsTable,
		userListTable,
	)
	var lists []domain.TodoList
	err := r.db.Select(&lists, query, userID)
	if err != nil {
		return nil, err
	}
	return &lists, nil
}

func (r *TodoListsRepository) GetListByID(listID int, userID int) (*domain.TodoList, error) {
	query := fmt.Sprintf(
		"select list.id, list.title, list.description "+
			"from %s list inner join %s user_list "+
			"on user_list.list_id = list.id "+
			"where list.id = $1 "+
			"and user_list.user_id = $2",
		todoListsTable,
		userListTable,
	)
	var list domain.TodoList
	err := r.db.Get(&list, query, listID, userID)
	if err != nil {
		return nil, err
	}
	return &list, nil
}
