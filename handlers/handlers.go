package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type HandlerConfig struct {
	Templates map[string]*template.Template
}

func (this *HandlerConfig) Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler")
	tmplName := getTemplateNameFromUrlPath(r.URL.Path)
	tmpl, ok := this.Templates[tmplName]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		log.Fatalf("Template %s not found: r.URL.Path %s", tmplName, r.URL.Path)
	}
	err := tmpl.ExecuteTemplate(w, tmplName+".html", nil)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}
}

func getTemplateNameFromUrlPath(urlPath string) string {
	if urlPath == "/" {
		return "home"
	} else {
		return urlPath[1:]
	}
}
