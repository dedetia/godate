package custerr

import (
	"errors"
	"net/http"
)

type HttpError struct {
	Code int
	Err  error
}

func (a *HttpError) Error() string {
	return a.Err.Error()
}

func New(code int, err error) *HttpError {
	var re *HttpError
	if errors.As(err, &re) {
		return re
	}
	return &HttpError{
		Code: code,
		Err:  err,
	}
}

func BadRequest(err error) *HttpError {
	return New(http.StatusBadRequest, err)
}

func Conflict(err error) *HttpError {
	return New(http.StatusConflict, err)
}

func Unauthorized(err error) *HttpError {
	return New(http.StatusUnauthorized, err)
}

func NotFound(err error) *HttpError {
	return New(http.StatusNotFound, err)
}

func Forbidden(err error) *HttpError {
	return New(http.StatusForbidden, err)
}
