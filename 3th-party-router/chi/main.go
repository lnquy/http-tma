package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// An HTTP server using chi router instead of default http.ServeMux.
// chi also supports grouping handlers, URL params, common middlewares and middlewares chaining...
//
// Things I love about chi is it's really lightweight, simple and 100% compatible with Go native net/http,
// which means you don't have to rewrite your code when switch to chi router.
// This is my library of choice when I have to write Go web/api server.
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer, middleware.DefaultCompress)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Index page with chi router"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Route("/users", func(r chi.Router) {
			r.Post("/", handlerPlaceHolder)
			r.Get("/", handlerPlaceHolder)
			r.Get("/{uid}", handlerPlaceHolder)
			r.Delete("/{uid}", handlerPlaceHolder)
			// ...
		})

		r.Route("/schools", func(r chi.Router) {
			r.Use(middleware.Timeout(10 * time.Second))
			r.Post("/", handlerPlaceHolder)
			r.Delete("/{sid}", handlerPlaceHolder)
			// ...
		})
	})

	r.Get("/login", handlerPlaceHolder)
	r.Post("/login", handlerPlaceHolder)
	r.Post("/logout", handlerPlaceHolder)

	http.ListenAndServe(":8080", r)
}

func handlerPlaceHolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" > Hit", r.URL.Path)
}
