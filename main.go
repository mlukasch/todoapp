package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"templatetest22/handlers"

	_ "github.com/lib/pq"
)

func main() {
	db := initDB()
	defer db.Close()

	handlers.Register(loadTemplates(), db)
	http.ListenAndServe(":8080", nil)
}

func initDB() *sql.DB {
	dbHost, _ := os.LookupEnv("POSTGRES_HOST")
	dbName, _ := os.LookupEnv("POSTGRES_DB")
	userName, _ := os.LookupEnv("POSTGRES_USER")
	password, _ := os.LookupEnv("POSTGRES_PASSWORD")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, userName, password, dbName))
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("ping db ok")
	return db
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
			tmpl, err := loadPageTemplates(dirName).ParseGlob("templates/*.gohtml")
			if err != nil {
				log.Fatal(err)
			}
			result[dirName] = tmpl
		}
	}
	for k, v := range result {
		log.Printf("Loaded Templates: %s mit: %#v", k, v.DefinedTemplates())
	}

	return result
}

func loadPageTemplates(dir string) *template.Template {
	tmpl, err := template.ParseGlob(path.Join("templates", dir, "*.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}
