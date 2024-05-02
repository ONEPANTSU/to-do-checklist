package domain

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

func (list *TodoList) ConvertFromArray(fields []interface{}) {
	id, _ := fields[0].(*int)
	title, _ := fields[1].(*string)
	description, _ := fields[2].(*string)
	list.ID = *id
	list.Title = *title
	list.Description = *description
}

func (list *TodoList) GetFields() []interface{} {
	return []interface{}{
		&list.ID,
		&list.Title,
		&list.Description,
	}
}

type UserList struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	ListID int `json:"list_id"`
}

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	ListID      int    `json:"list_id"`
}
