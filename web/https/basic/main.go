package main

import "net/http"

// A basic HTTPS server.
func main() {
	r := &http.ServeMux{}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("A web page on HTTPS"))
	})

	http.ListenAndServeTLS(":8080", "../cert/server.crt", "../cert/server.key", r)
}
