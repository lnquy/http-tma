package main

import (
	"html/template"
	"net/http"
)

// Basic web server that serves an HTML page with dynamic data binding using Go html/template.
// This is what equivalent to MVC model.
func main() {
	r := &http.ServeMux{}
	r.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", r)
}

// Controller
func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTmpl, err := template.ParseFiles("./tmpl/index.html") // View
	if err != nil {
		http.Error(w, "failed to parse index.html template", http.StatusInternalServerError)
		return
	}

	fakeData := IndexPageData{
		Title: "Index page with dynamic data binding",
		Users: []User{
			{123, "Edsger W. Dijkstra", "Dutch"},
			{456, "Alan Turing", "UK"},
			{789, "Pierre de Fermat", "France"},
		},
	}

	indexTmpl.Execute(w, fakeData)
}

// Model
type (
	IndexPageData struct {
		Title string
		Users []User
	}

	User struct {
		Id      int
		Name    string
		Country string
	}
)
