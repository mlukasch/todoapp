package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"templatetest22/handlers"
)

func main() {
	c := loadTemplates()
	handlerConfig := handlers.HandlerConfig{
		c,
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", handlerConfig.Handler)
	http.ListenAndServe(":8080", nil)
}

func loadTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	infos, err := ioutil.ReadDir("templates")
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if info.IsDir() {
			dirName := info.Name()
			result[dirName] = template.Must(loadPageTemplates(dirName).ParseGlob("templates/*.html"))
		}
	}
	return result
}

func loadPageTemplates(dir string) *template.Template {
	pageTmpl := template.Must(template.ParseGlob(path.Join("templates", dir, "*.html")))
	return pageTmpl
}
