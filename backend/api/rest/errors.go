package rest

import (
	"auto-parts-catalog/backend/catalogservice"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var (
	ErrNotFound            = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
	ErrBadRequest          = &ErrResponse{HTTPStatusCode: 400, StatusText: "Bad request"}
	ErrAlreadyExists       = &ErrResponse{HTTPStatusCode: 409, StatusText: "Already exists"}
	ErrInternalServerError = &ErrResponse{HTTPStatusCode: 500, StatusText: "Internal Server Error"}

	ErrBind = errors.New("Bind error")
)

func serviceToHttpError(err error) *ErrResponse {
	if err == nil {
		return nil
	}

	var storageErr *catalogservice.CatalogServiceError

	if errors.As(err, &storageErr) {
		switch storageErr.Code {
		case catalogservice.InvalidArgumentError:
			return ErrBadRequest
		case catalogservice.NotFoundError:
			return ErrNotFound
		case catalogservice.AlreadyExistsError:
			return ErrAlreadyExists
		default:
			return ErrInternalServerError
		}
	} else {
		return ErrInternalServerError
	}

}
