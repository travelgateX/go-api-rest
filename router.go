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

func newRouter() *chi.Mux {
	jwtParserConfig := jwt.ParserConfig{
		AdminGroup:       "admin",
		PublicKey:        "-----BEGIN CERTIFICATE-----\nMIIDAzCCAeugAwIBAgIJBNrNC7THKVOtMA0GCSqGSIb3DQEBCwUAMB8xHTAbBgNV\nBAMTFHh0Zy1kZXYuZXUuYXV0aDAuY29tMB4XDTE3MDkwNDE2MTIyMFoXDTMxMDUx\nNDE2MTIyMFowHzEdMBsGA1UEAxMUeHRnLWRldi5ldS5hdXRoMC5jb20wggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC6dh6RISMxNrOrlTyXpI3oW8M6H389\nz7bMybyXEgBq35dWTmLLbX/lbUzv/7S7AT0hN31F0v3Eyoxt20a1pjjJacbGH3px\ntbhEuYTCb9Jf+OLBThVGQtQiS/jtRHiTx5vW1w6oMEpaYTXqstSzCqyG7GqbH+rO\nVEg1EA1ISSPhDEB8btBlPdEf7auUvyLbxkzX/CIhekBbqsSMBIZYzWZy0ht56W8k\nZk34S23bLEfjIWydTERATFi+QmCsHnIWclLh2aFzl/cvLxBZ+D/Ayy2AdoAB7wIY\nK011O31DSNrY23GBhUMMMWZZwg9IY1oVXIyheeoFebHPmNNU7yQkiUInAgMBAAGj\nQjBAMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFFqpUMSQoxOfy+/q9Rgbulst\n6XwRMA4GA1UdDwEB/wQEAwIChDANBgkqhkiG9w0BAQsFAAOCAQEAOoXal9KDjAUm\nnOkKg44TDd7qz/yXJgmra/G7jZD1Y7w/SCfFps9gR0PCrO5oYAoEwjetSIfGb/52\nkN8uN9MEiyBzTrrbVO16EzWk4Mw69aFEY3SeWFEXtevmWCM7OmcQcf+6IwMkk2BI\nVkh14M5Ybkj+IYtGlWiceJj7GeqHGws02dzYNQf3hBQSGJ891bSgx+C4H7Maxd6B\nq69JuYQhl7bSetlQ0mxLRTSD2mGyqetfaGuBb89LwGApokzlWyYKStKoJivss8yB\n0sa6/PtevO2+RHr/QacrA7MoATIZi8fuX7bCrCpulWCd9bELwlSEjkW5dD2isFN6\nAjXpFHDQ8Q==\n-----END CERTIFICATE-----",
		DummyToken:       "dummyToken",
		IgnoreExpiration: false,
		GroupsClaim:      []string{"https://xtg.com/iam", "https://travelgatex.com/iam"},
		MemberIDClaim:    []string{"https://xtg.com/member_id", "https://travelgatex.com/member_id"},
	}

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
		authMiddleware,
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
