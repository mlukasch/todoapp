package db

import (
	"database/sql"
)

func InsertTodoUser(db *sql.DB, user *TodoUser) error {
	_, err := db.Exec("INSERT INTO TODO_USER (firstName,lastName,email,password) VALUES ($1,$2,$3,$4)", user.FirstName, user.LastName, user.Email, user.Password)
	return err
}
