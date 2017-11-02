package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

type HandlerConfig struct {
	templates map[string]*template.Template
}

func Register(templates map[string]*template.Template) {
	config := HandlerConfig{
		templates,
	}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/home", config.HomeHandler)
	http.HandleFunc("/about", config.AboutHandler)
	http.HandleFunc("/todos/new", config.TodoFormHandler)
	http.HandleFunc("/todos", config.TodoListHandler)
}

func (this *HandlerConfig) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler")
	tmpl := this.templates["home"]
	err := tmpl.Execute(w, nil)
	if err != nil {
		HandleErrorIfExists(err, w)
	}
}

func (this *HandlerConfig) AboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AboutHandler")
	tmpl := this.templates["about"]
	err := tmpl.Execute(w, nil)
	if err != nil {
		HandleErrorIfExists(err, w)
	}
}

func (this *HandlerConfig) TodoFormHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TodoFormHandler")
	tmpl := this.templates["todo_form"]
	err := tmpl.Execute(w, nil)
	if err != nil {
		HandleErrorIfExists(err, w)
	}
}

func (this *HandlerConfig) TodoListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TodoListHandler")
	tmpl := this.templates["todo_list"]
	err := tmpl.Execute(w, nil)
	if err != nil {
		HandleErrorIfExists(err, w)
	}
}

func (this *HandlerConfig) Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler")
	tmplName := getTemplateNameFromUrlPath(r.URL.Path)
	tmpl, ok := this.templates[tmplName]
	if !ok {
		HandleErrorIfExists(errors.New("Template not found"), w)
		return
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		HandleErrorIfExists(err, w)
	}
}

func getTemplateNameFromUrlPath(urlPath string) string {
	if urlPath == "/" || len(urlPath) == 0 {
		return "home"
	} else {
		return urlPath[1:]
	}
}

func HandleErrorIfExists(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
	log.Printf("Error : %s", err.Error())
}
