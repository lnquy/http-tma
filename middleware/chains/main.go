package main

import (
	"fmt"
	"net/http"
)

// Chaining multiple middlewares.
func main() {
	r := &http.ServeMux{} // Router
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/panic", func(writer http.ResponseWriter, r *http.Request) {
		panic("stupid panic")
	})

	http.ListenAndServe(":8080", chainMiddlewares(r, recoveryMiddleware, logMiddleware))
}

// Handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// Middleware
type Middleware func(http.Handler) http.HandlerFunc

func logMiddleware(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(" > Request hit: %s\n", r.URL.Path)
		h.ServeHTTP(w, r) // Pass request to the next func (middleware/handler)
	})
}

func recoveryMiddleware(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf(" > PANIC (but recovered): %s", r)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func chainMiddlewares(h http.Handler, mdws ...Middleware) http.Handler {
	for _, m := range mdws {
		h = m(h)
	}
	return h
}
