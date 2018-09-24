package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// An HTTP server using gorilla/mux router instead of default http.ServeMux.
// mux is basically the same as chi, lightweight, simple and 100% compatible with Go native net/http.
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerPlaceHolder)
	r.HandleFunc("/products", handlerPlaceHolder).Methods("POST")
	r.HandleFunc("/articles", handlerPlaceHolder).Methods("GET")
	r.HandleFunc("/articles/{id}", handlerPlaceHolder).Methods("GET", "PUT")
	r.HandleFunc("/authors", handlerPlaceHolder).Queries("surname", "{surname}")

	http.Handle("/", r)
}

func handlerPlaceHolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" > Hit", r.URL.Path)
}
