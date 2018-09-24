package main

import (
	"encoding/json"
	"net/http"
)

// Basic HTTP server that serves a JSON API.
func main() {
	r := &http.ServeMux{}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Msg string `json:"msg"`
		}{"Hello, this is a JSON message!"}

		b, _ := json.Marshal(data) // Encode data object to JSON's []byte
		w.Write(b)                 // Response JSON []byte to client
	})

	http.ListenAndServe(":8080", r)
}
