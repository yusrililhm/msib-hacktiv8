package errs

import "net/http"

type Error interface {
	Message() string
	Status() int
	Error() string
}

type ErrorData struct {
	ErrMessage string
	ErrStatus  int
	ErrError   string
}

// Error implements Error.
func (e *ErrorData) Error() string {
	return e.ErrError
}

// Message implements Error.
func (e *ErrorData) Message() string {
	return e.ErrMessage
}

// Status implements Error.
func (e *ErrorData) Status() int {
	return e.ErrStatus
}

func NewUnauthorizedError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   "NOT_AUTHORIZED",
	}
}

func NewUnauthenticatedError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "NOT_AUTHENTICATED",
	}
}

func NewNotFoundError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "NOT_FOUND",
	}
}

func NewBadRequestError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessableEntityError(message string) Error {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "INVALID_REQUEST_BODY",
	}
}
