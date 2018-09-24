package main

import (
	"fmt"
	"net/http"
)

// Custom middleware that log the response's status to the console before returning back to client.
func main() {
	r := &http.ServeMux{}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root"))
	})
	r.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "some server error", http.StatusInternalServerError)
		w.Write([]byte("500"))
	})

	http.ListenAndServe(":8080", logResponseStatusMiddleware(r))
}

// Wrap http.ResponseWriter so we can do other logic before response back to client
type myResponseWriter struct {
	// Embedded, so myResponseWriter already "inherited" all methods from http.ResponseWriter.
	// => myResponseWriter now can be used as http.ResponseWriter.
	http.ResponseWriter
	status int // Track the final response status
}

// Override the WriteHeader method of http.ResponseWriter, so we can track the response status.
func (m *myResponseWriter) WriteHeader(s int) {
	m.status = s
	m.ResponseWriter.WriteHeader(s)
}

func logResponseStatusMiddleware(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myRW := &myResponseWriter{w, 200}
		h.ServeHTTP(myRW, r)

		// This line of code will be called when handler already finished, so we can log the final response status,
		// before responding back to client.
		fmt.Printf(" > Request (%s) has status: %d\n", r.URL.Path, myRW.status)
	})
}
