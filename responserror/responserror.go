package responserror

import (
	"net/http"

	"github.com/go-chi/render"
)

// ResponseError : error main struct
type ResponseError struct {
	Err error `json:"-"` // low-level runtime error

	StatusCode int    `json:"status"`
	StatusText string `json:"description"` // user-level status message
	ErrorText  string `json:"error"`       // application-level error message, for debugging
}

// Render : render error
func (a *ResponseError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, a.StatusCode)
	return nil
}

// ErrBadRequest : Bad request pre-formatted response
func ErrBadRequest(err error) render.Renderer {
	return &ResponseError{
		Err:        err,
		StatusCode: 400,
		StatusText: "Invalid request",
		ErrorText:  err.Error(),
	}
}

// ErrInternal : internal error pre-formatted response
func ErrInternal(err error) render.Renderer {
	return &ResponseError{
		Err:        err,
		StatusCode: 500,
		StatusText: "Internal error",
		ErrorText:  err.Error(),
	}
}

// ErrResourceNotFound : resource not found pre-formatted response
func ErrResourceNotFound() render.Renderer {
	return &ResponseError{
		StatusCode: 404,
		StatusText: "Resource not found.",
	}
}
