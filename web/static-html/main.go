package main

import (
	"net/http"
)

// A basic HTTP server that serves static html files from html folder.
func main() {
	// Router
	r := &http.ServeMux{}
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./html")))) // Bind html folder to /static URL

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
}
