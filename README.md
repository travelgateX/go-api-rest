# go-api-rest 

This project is a base structure API for making POST and GET requests. It's meant to serve as a reference to build your own REST API. 

# Installation
For this project, first you will need to install the *go-chi* package.
```bash
go get -u github.com/go-chi/chi	
```
You also will need the render subpackage.
```bash
go get github.com/go-chi/render
```

sqlx library 
```bash  
go get github.com/jmoiron/sqlx
```

# Add your package
There is an example package <code>todo</code>. You can add your own package with your functions. In order to do it:

## Middleware
At route.go you have the middleware declaration.

```go
func newRouter() *chi.Mux {
	// New chi router
	router := chi.NewRouter()

	// Api middlewares
	router.Use(
	render.SetContentType(render.ContentTypeJSON),
		middleware.Logger, // remove in production
		middleware.DefaultCompress,
		middleware.Recoverer,
		middleware.YourMiddleware
	)
        ...

	return router
}
```

## Routes and functions
Create a directory at the same level as <code>todo</code> package.
Then add your rotues following the pattern: 
```go
// Routes :
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetATodo)
	router.Delete("/{id}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)
	return router
}
```

Then you must add it and mount it to the version you want:

```go

	// Api root route "/v1"
	router.Route("/v1", func(r chi.Router) {
        r.Mount("/api/todo", todo.Routes())
        r.Mount("api/YOUPACKAGE",YOURPACKAGE.Routes())
	})
```

That should be enough to have a small functional REST API.
