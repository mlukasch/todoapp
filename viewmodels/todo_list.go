package viewmodels

import (
	"time"
)

type TodoListItem struct {
	title        string
	creator      User
	creationDate time.Time
}

type User struct {
	firstName string
	lastName  string
	email     string
}

type TodoList struct {
	todoList []TodoListItem
	user     User
}

func GetSample() TodoList {
	user := User{
		firstName: "Mr",
		lastName:  "L",
		email:     "bdbdfgdllas@adad.de",
	}

	items := []TodoListItem{
		TodoListItem{
			"Einkaufen",
			user,
			time.Now().Add(48*time.Hour - 10*time.Minute),
		},
		TodoListItem{
			"Frühstücken",
			user,
			time.Now().Add(48*time.Hour - 10*time.Minute),
		},
		TodoListItem{
			"Pommes essen",
			user,
			time.Now().Add(48*time.Hour - 10*time.Minute),
		},
	}
	return TodoList{
		items,
		user,
	}
}
