package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

var (
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrNotFound            = errors.New("not found")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrInternalServerError = errors.New("internal server error")
)

// HttpErr interface
type HttpErr interface {
	Status() int
	Error() string
	Details() interface{}
}

// HttpError struct
type HttpError struct {
	ErrStatus  int         `json:"status"`
	ErrError   string      `json:"error"`
	ErrDetails interface{} `json:"details"`
}

// Error  Error() interface method
func (e HttpError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - details: %v", e.ErrStatus, e.ErrError, e.ErrDetails)
}

// Error status
func (e HttpError) Status() int {
	return e.ErrStatus
}

// HttpError Details
func (e HttpError) Details() interface{} {
	return e.ErrDetails
}

// New Http Error
func NewHttpError(status int, err string, details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  status,
		ErrError:   err,
		ErrDetails: details,
	}
}

// New Bad Request Error
func NewBadRequestError(details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  http.StatusBadRequest,
		ErrError:   ErrBadRequest.Error(),
		ErrDetails: details,
	}
}

// New Unauthorized Error
func NewUnauthorizedError(details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   ErrUnauthorized.Error(),
		ErrDetails: details,
	}
}

// New Forbidden Error
func NewForbiddenError(details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  http.StatusForbidden,
		ErrError:   ErrForbidden.Error(),
		ErrDetails: details,
	}
}

// New Not Found Error
func NewNotFoundError(details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  http.StatusNotFound,
		ErrError:   ErrNotFound.Error(),
		ErrDetails: details,
	}
}

// New Unprocessable Entity Error
func NewUnprocessableEntityError(details interface{}) HttpErr {
	return HttpError{
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   ErrUnprocessableEntity.Error(),
		ErrDetails: details,
	}
}

// New Internal Server Error
func NewInternalServerError(details interface{}) HttpErr {
	log.Error(details.(error).Error())

	return HttpError{
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   ErrInternalServerError.Error(),
		ErrDetails: nil,
	}
}

// Parse Http Error
func ParseHttpError(err error) (int, interface{}) {
	if httpErr, ok := err.(HttpErr); ok {
		return httpErr.Status(), httpErr
	}
	return http.StatusInternalServerError, NewInternalServerError(err)
}
