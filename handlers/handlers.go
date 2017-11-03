package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"templatetest22/utils"
)

type HandlerConfig struct {
	templates map[string]*template.Template
}

func Register(templates map[string]*template.Template) {
	config := HandlerConfig{
		templates,
	}

	http.HandleFunc("/", config.HomeHandler)
	http.HandleFunc("/about", config.AboutHandler)
	http.HandleFunc("/todos/new", config.TodoFormHandler)
	http.HandleFunc("/todos", config.TodoListHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
}

func (this *HandlerConfig) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler")
	if r.Method == http.MethodPost {
		log.Println("HomeHandler Post")
		userName := r.FormValue("userName")
		email := r.FormValue("email")
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
