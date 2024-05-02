package domain

type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required" db:"username"`
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"hashed_password"`
}

func (user *User) ConvertFromArray(fields []interface{}) {
	id, _ := fields[0].(*int)
	username, _ := fields[1].(*string)
	email, _ := fields[2].(*string)
	password, _ := fields[3].(*string)
	user.ID = *id
	user.Username = *username
	user.Email = *email
	user.Password = *password
}

func (user *User) GetFields() []interface{} {
	return []interface{}{
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	}
}

type SignIn struct {
	Username string
	Password string
}
