package main

import (
	"net/http"
)

// A basic HTTP server that serves 3 APIs at: /, /hello and /api.
func main() {
	// Router
	r := &http.ServeMux{}
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API"))
	})
	r.Handle("/hello", &helloHandler{})

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
	// l, _ := net.Listen("tcp", ":8080")
	// http.Serve(l, r)
}

// Handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm Groot!"))
}

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
