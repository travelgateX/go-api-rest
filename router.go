package main

import (
	"billing-calculation-center/todo"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func newRouter() *chi.Mux {
	// New chi router
	router := chi.NewRouter()

	// Api middlewares
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger, // remove in production
		middleware.DefaultCompress,
		middleware.Recoverer,
	)

	// Api root route "/v1"
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todo.Routes())
	})

	// Log all API routes & middlewares
	walkFunc := func(method string, route string, handlder http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err : %s\n", err.Error()) // panic if there is an error
	}

	return router
}
