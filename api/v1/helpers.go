package v1

import (
	"net/http"
)

func BadRequest(err error) (int, error) {
	return http.StatusBadRequest, err
}

func InternalServerError(err error) (int, error) {
	return http.StatusInternalServerError, err
}

func NotFound(err error) (int, error) {
	return http.StatusNotFound, err
}

func OK() (int, error) {
	return http.StatusOK, nil
}

func Created() (int, error) {
	return http.StatusCreated, nil
}

func NoContent() (int, error) {
	return http.StatusNoContent, nil
}
