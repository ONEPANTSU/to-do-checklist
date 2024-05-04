package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"to-do-checklist/internal/domain"
)

type TodoItemRepository struct {
	db *sqlx.DB
}

func newTodoItemRepository(db *sqlx.DB) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

func (r *TodoItemRepository) CreateItem(item *domain.TodoItem) (int, error) {
	query := fmt.Sprintf(
		"insert into %s (title, description, completed, list_id) "+
			"values ($1, $2, $3, $4) "+
			"returning id",
		todoItemsTable,
	)
	row := r.db.QueryRow(query, item.Title, item.Description, item.Completed, item.ListID)
	var itemID int
	if err := row.Scan(&itemID); err != nil {
		return 0, err
	}
	return itemID, nil
}

func (r *TodoItemRepository) GetItems(listID int) (*[]domain.TodoItem, error) {
	query := fmt.Sprintf("select * from %s where list_id = $1", todoItemsTable)
	var items []domain.TodoItem
	err := r.db.Select(&items, query, listID)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (r *TodoItemRepository) GetItemByID(itemID, userID int) (*domain.TodoItem, error) {
	var item domain.TodoItem
	query := fmt.Sprintf(
		"select items.id, items.title, items.description, items.completed, items.list_id "+
			"from %s items inner join %s users_lists on items.list_id = users_lists.list_id "+
			"where users_lists.user_id = $1 and items.id = $2",
		todoItemsTable,
		userListTable,
	)
	if err := r.db.Get(&item, query, userID, itemID); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *TodoItemRepository) UpdateItem(item *domain.UpdateTodoItem, itemID int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if item.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, item.Title)
		argID++
	}
	if item.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, item.Description)
		argID++
	}
	if item.Completed != nil {
		setValues = append(setValues, fmt.Sprintf("completed=$%d", argID))
		args = append(args, item.Completed)
		argID++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(
		"update %s items set %s "+
			"where items.id = $%d",
		todoItemsTable,
		setQuery,
		argID,
	)
	fmt.Println(query)
	args = append(args, itemID)
	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}
	return nil
}

func (r *TodoItemRepository) DeleteItem(itemID int) error {
	query := fmt.Sprintf("delete from %s where id = $1", todoItemsTable)
	if _, err := r.db.Exec(query, itemID); err != nil {
		return err
	}
	return nil
}
