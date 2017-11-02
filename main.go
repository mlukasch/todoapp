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
	handlers.Register(loadTemplates())
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
			result[dirName] = template.Must(loadPageTemplates(dirName).ParseGlob("templates/*.gohtml"))
		}
	}
	for k, v := range result {
		log.Printf("Loaded Templates: %s mit: %#v", k, v.DefinedTemplates())
	}

	return result
}

func loadPageTemplates(dir string) *template.Template {
	pageTmpl := template.Must(template.ParseGlob(path.Join("templates", dir, "*.gohtml")))
	return pageTmpl
}
