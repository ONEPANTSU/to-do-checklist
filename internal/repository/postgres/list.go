package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (r *TodoListsRepository) GetListByID(listID, userID int) (*domain.TodoList, error) {

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

func (r *TodoListsRepository) UpdateList(list *domain.UpdateTodoList, listID, userID int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if list.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, list.Title)
		argID++
	}
	if list.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, list.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"update %s list set %s "+
			"from %s user_list "+
			"where user_list.list_id = list.id "+
			"and list.id = $%d "+
			"and user_list.user_id = $%d",
		todoListsTable,
		setQuery,
		userListTable,
		argID,
		argID+1,
	)
	args = append(args, listID, userID)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoListsRepository) DeleteList(listID, userID int) error {
	query := fmt.Sprintf(
		"delete from %s list using %s user_list "+
			"where user_list.list_id = list.id and "+
			"list_id = $1 and user_list.user_id = $2",
		todoListsTable,
		userListTable,
	)
	_, err := r.db.Exec(query, listID, userID)
	if err != nil {
		return err
	}
	return nil
}
