package handlers

import (
	"database/sql"
	"errors"
	"html"
	"html/template"
	"log"
	"net/http"

	"templatetest22/db"
	"templatetest22/utils"
	"templatetest22/validators"
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
	http.HandleFunc("/about", config.AboutHandler)
	http.HandleFunc("/todos/new", config.TodoFormHandler)
	http.HandleFunc("/user/new", config.UserNew)
	http.HandleFunc("/todos", config.TodoListHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
}

func (this *HandlerConfig) UserNew(w http.ResponseWriter, r *http.Request) {
	tmpl := this.templates["home"]

	// Validation:
	validator := validators.NewForm()
	validator.AddField("firstName", r.FormValue("firstName")).Add(validators.NotEmptyValidator)
	validator.AddField("lastName", r.FormValue("lastName")).Add(validators.NotEmptyValidator)
	validator.AddField("email", r.FormValue("email")).Add(validators.NotEmptyValidator)
	validator.AddField("password", r.FormValue("password")).Add(validators.NotEmptyValidator)
	ok, errors := validator.Execute()

	if len(errors) > 0 {
		tmpl.ExecuteTemplate(w, "home.gohtml", struct {
			HasErrors bool
			Errors    map[string][]error
		}{
			!ok,
			errors,
		})
		return
	}

	todoUser := db.TodoUser{
		FirstName: html.EscapeString(r.FormValue("firstName")),
		LastName:  html.EscapeString(r.FormValue("lastName")),
		Email:     html.EscapeString(r.FormValue("email")),
		Password:  r.FormValue("password"),
	}
	_, err := db.InsertTodoUser(this.db, &todoUser)
	if err != nil {
		log.Println(err)
	}
	//successful := (err == nil)

	err = tmpl.ExecuteTemplate(w, "home.gohtml", struct {
		Registered bool
		FirstName  string
	}{
		true,
		todoUser.FirstName,
	})
	if err != nil {
		utils.HandleAsNotFound(err, w)
	}
}

func (this *HandlerConfig) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler")

	log.Println("HomeHandler Post")
	userName := r.FormValue("userName")
	email := r.FormValue("email")
	// TODO add validation
	log.Println(userName)
	log.Println(email)
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
