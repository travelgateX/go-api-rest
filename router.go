package main

import (
	"log"
	"net/http"

	"go-api-rest/auth"
	"go-api-rest/todo"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	authorization "github.com/travelgateX/go-jwt-tools"
	"github.com/travelgateX/go-jwt-tools/jwt"
)

func newRouter(jwtParserConfig jwt.ParserConfig) *chi.Mux {
	jwtParser := jwt.NewParser(jwtParserConfig)
	authMiddleware := authorization.Middleware(jwtParser)

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

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware)
			r.Get("/api/todo/jwt_example/", JwtExample)
		})
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

// JwtExample : Return all valid groups in jwt token (example)
func JwtExample(w http.ResponseWriter, r *http.Request) {
	a := auth.NewAuth(r.Context())
	validGroups := a.CheckAllPermission("iam", "grp", "r")
	render.JSON(w, r, validGroups)
}
