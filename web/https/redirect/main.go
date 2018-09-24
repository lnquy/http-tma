package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Start a child HTTP server runs on port 80 which redirect all request to HTTPS server on port 8443
	go func() {
		fmt.Println(http.ListenAndServe(":80", http.HandlerFunc(redirectHTTPToHTTPS)))
	}()

	// HTTPS server
	r := &http.ServeMux{}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("A web page on HTTPS"))
	})
	fmt.Println(http.ListenAndServeTLS(":8443", "../cert/server.crt", "../cert/server.key", r))
}

// Middleware
func redirectHTTPToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+":8443"+r.URL.Path, http.StatusTemporaryRedirect)
}
