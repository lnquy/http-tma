package main

import (
	"fmt"
	"net/http"
)

// A basic HTTP server that serves 3 APIs: /, /hello and /api.
// Also log all request to console by a global middleware.
func main() {
	// Router
	r := &http.ServeMux{}
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API"))
	})
	r.Handle("/hello", &helloHandler{})

	// Start the HTTP server
	http.ListenAndServe(":8080", logMiddleware(r))
}

// Handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm Groot!"))
}

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// Middleware
func logMiddleware(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("   > HIT: %s (%s)\n", r.URL.Path, r.RemoteAddr)
		h.ServeHTTP(w, r)
		fmt.Printf("   > END: %s (%s)\n", r.URL.Path, r.RemoteAddr)
	})
}
