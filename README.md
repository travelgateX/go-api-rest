# go-api-rest 

This project is a base structure API for making POST and GET requests. It's meant to serve as a reference to build your own REST API. 

# Installation
This project uses Go Modules and requires Go 1.11+. If your Go version is lesser than 1.11, install the dependeces inside the go.mod file.
```bash
git clone https://github.com/travelgateX/go-api-rest
```
# Configuration
## Database
To add your database configuration take the **config.example** file and modify the [db] part with your data. Then rename the file to config.toml
```toml
[db]
host = "YOUR-HOST"
port = "5432"
user = "USER"
pass = "PASSWORD"
name = "DBNAME"
ssl  = false
 ```

If you want to try this exact example, run this command on your database.
```sql
CREATE TABLE public.todo (
    id SERIAL PRIMARY KEY,
    title character varying(255),
    body character varying(255)
);
```

## Authorization
If you want to test the authorization middleware, take the **config.example** file and modify the [auth] part with your authorization data. 


There is a specific route to test the authorization package: **/v1/api/todo/jwt_example**. 

*Note: If the config file is left empty, all the other querys will work anyway.*
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
Then you just need to run the projetc!
```bash
go run .
```
That should be enough to have a small functional REST API.