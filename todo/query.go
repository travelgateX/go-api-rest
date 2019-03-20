package todo

import (
	"errors"

	"go-api-rest/sql"
)

// Get a todo
func getATodo(id string) (*[]Todo, error) {
	var todo []Todo

	query := "SELECT * FROM public.todo WHERE id = " + stringToSQL(&id)
	if err := sql.Instance.DB.Select(&todo, query); err != nil {
		return nil, err
	}

	if len(todo) == 0 {
		return nil, errors.New("Todo not found")
	}

	return &todo, nil
}

// Get all todos
func getAllTodos() ([]Todo, error) {
	query := "SELECT * FROM public.todo"
	var allTodos []Todo
	errSQL := sql.Instance.DB.Select(&allTodos, query)
	if errSQL != nil {
		return allTodos, errSQL
	}
	return allTodos, nil
}

func createTodo(input Todo) (*Todo, error) {
	mutation := "INSERT INTO public.todo(title,body) values (" + stringToSQL(&input.Title) + "," + stringToSQL(&input.Body) + ")" + " returning id,body,title"
	var todo Todo
	errSQL := sql.Instance.DB.QueryRow(mutation).Scan(&todo.ID, &todo.Body, &todo.Title)
	if errSQL != nil {
		return nil, errSQL
	}
	return &todo, nil
}

func deleteTodo(id string) (*Todo, error) {
	// Check if todo exists in database
	var todoID []Todo
	query := "SELECT * FROM public.todo WHERE id = " + stringToSQL(&id)
	errSQL := sql.Instance.DB.Select(&todoID, query)
	if errSQL != nil {
		return nil, errSQL
	}
	if len(todoID) == 0 {
		err := errors.New("Todo not found")
		return nil, err
	}
	// If exists, delete it from database
	mutation := "DELETE FROM public.todo WHERE id = " + stringToSQL(&id)
	r, errSQL := sql.Instance.DB.Queryx(mutation)
	defer r.Close()
	if errSQL != nil {
		return nil, errSQL
	}

	return &todoID[0], nil
}

func stringToSQL(value *string) string {
	if value == nil {
		return ""
	}
	result := "'" + *value + "'"
	return result
}
