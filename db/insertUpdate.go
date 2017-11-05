package db

import (
	"database/sql"
)

func InsertTodoUser(db *sql.DB, user *TodoUser) (int64, error) {
	result, err := db.Exec("INSERT INTO TODO_USER (firstName,lastName,email,password) VALUES ($1,$2,$3,$4)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
