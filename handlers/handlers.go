package handlers

import (
	"database/sql"
	"errors"
	"html/template"
	"log"
	"net/http"

	"templatetest22/db"
	"templatetest22/utils"
)

type HandlerConfig struct {
	templates map[string]*template.Template
	db        *sql.DB
}

func Register(templates map[string]*template.Template, db *sql.DB) {
	config := HandlerConfig{
		templates,
		db,
	}

	http.HandleFunc("/", config.HomeHandler)
	http.HandleFunc("/register", config.RegisterHandler)
	http.HandleFunc("/about", config.AboutHandler)
	http.HandleFunc("/todos/new", config.TodoFormHandler)
	http.HandleFunc("/todos", config.TodoListHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
}

func (this *HandlerConfig) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RegisterHandler")
	if r.Method == http.MethodPost {

		todoUser := db.TodoUser{
			FirstName: r.FormValue("firstName"),
			LastName:  r.FormValue("lastName"),
			Email:     r.FormValue("email"),
			Password:  r.FormValue("password"),
		}
		err := db.InsertTodoUser(this.db, &todoUser)
		successful := (err != nil)
		tmpl := this.templates["home"]
		err = tmpl.ExecuteTemplate(w, "home.gohtml", struct{ Registered bool }{
			successful,
		})
		if err != nil {
			utils.HandleAsNotFound(err, w)
		}
	}
}

func (this *HandlerConfig) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler")
	if r.Method == http.MethodPost {
		log.Println("HomeHandler Post")
		userName := r.FormValue("userName")
		email := r.FormValue("email")
		// TODO add validation
		log.Println(userName)
		log.Println(email)
	}

	tmpl := this.templates["home"]
	err := tmpl.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		utils.HandleAsNotFound(err, w)
	}
}

func (this *HandlerConfig) AboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AboutHandler")
	tmpl, ok := this.templates["about"]
	if !ok {
		utils.HandleAsNotFound(errors.New("Template not found"), w)
		return
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		utils.HandleAsNotFound(err, w)
	}
}

func (this *HandlerConfig) TodoFormHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TodoFormHandler")
	tmpl, ok := this.templates["todo_form"]
	if !ok {
		utils.HandleAsNotFound(errors.New("Template not found"), w)
		return
	}
	err := tmpl.ExecuteTemplate(w, "todo_form.gohtml", nil)
	if err != nil {
		utils.HandleAsNotFound(err, w)
	}
}

func (this *HandlerConfig) TodoListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TodoListHandler")
	tmpl, ok := this.templates["todo_list"]
	if !ok {
		utils.HandleAsNotFound(errors.New("Template not found"), w)
		return
	}
	err := tmpl.ExecuteTemplate(w, "todo_list.gohtml", nil)
	if err != nil {
		utils.HandleAsNotFound(err, w)
	}
}

func (this *HandlerConfig) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleAsNotFound(errors.New("Resource not found"), w)
}

func getTemplateNameFromUrlPath(urlPath string) string {
	if urlPath == "/" || len(urlPath) == 0 {
		return "home"
	} else {
		return urlPath[1:]
	}
}
