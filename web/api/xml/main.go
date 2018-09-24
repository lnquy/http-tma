package main

import (
	"encoding/xml"
	"net/http"
)

// Basic HTTP server that serves a XML API.
func main() {
	r := &http.ServeMux{}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Msg string `json:"msg"`
		}{"Hello, this is a XML message!"}

		b, _ := xml.Marshal(data) // Encode data object to XML's []byte
		w.Write(b)                 // Response XML []byte to client
	})

	http.ListenAndServe(":8080", r)
}
