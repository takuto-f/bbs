package main

import (
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Content []byte
}

var templates = template.Must(template.ParseFiles("templates/index.tmpl.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := &Page{Content: []byte("Hello, World")}
	err := templates.ExecuteTemplate(w, "index.tmpl.html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, nil)
}
