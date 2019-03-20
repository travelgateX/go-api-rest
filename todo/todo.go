package todo

import (
	"errors"
	"net/http"

	"go-api-rest/responserror"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Todo :
type Todo struct {
	ID    string `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
}

var list []Todo

// Routes :
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetATodo)
	router.Delete("/{id}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)
	return router
}

// GetATodo : Get a todo by ID
func GetATodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp, err := getATodo(id)
	if err != nil {
		render.JSON(w, r, responserror.ErrBadRequest(err))
	} else {
		render.JSON(w, r, resp)
	}
}

// DeleteTodo : Delete a todo by ID
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp, err := deleteTodo(id)
	if err != nil {
		render.JSON(w, r, responserror.ErrBadRequest(err))
	} else {
		render.JSON(w, r, resp)
	}
}

// CreateTodo : Create a todo using post request
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todo := Todo{Body: r.FormValue("body"), Title: r.FormValue("title")}
	if formErrors := todo.validateBody(); formErrors != nil {
		render.JSON(w, r, responserror.ErrBadRequest(formErrors))
	} else {
		resp, err := createTodo(todo)
		if err != nil {
			render.JSON(w, r, responserror.ErrBadRequest(err))
		} else {
			render.JSON(w, r, resp)
		}
	}
}

// GetAllTodos : Return all todo stored in database
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	allTodos, err := getAllTodos()
	if err != nil {
		render.JSON(w, r, responserror.ErrInternal(err))
	} else {
		render.JSON(w, r, allTodos)
	}
}

func (t *Todo) validateBody() error {
	if t.Body == "" {
		return errors.New("body field is required")
	}

	if t.Title == "" {
		return errors.New("title field is required")
	}
	return nil
}
