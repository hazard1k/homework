package v1

import (
	"encoding/json"
	jsonapi2 "goarch/app/presentors/jsonapi"
	"net/http"
)

func JsonError(w http.ResponseWriter, status int, err error) {
	j, _ := json.Marshal(jsonapi2.NewErrorResponse(status, err))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

func JsonSuccess(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func BadRequest(err error) (int, []byte, error) {
	return http.StatusBadRequest, nil, err
}

func InternalServerError(err error) (int, []byte, error) {
	return http.StatusInternalServerError, nil, err
}

func NotFound(err error) (int, []byte, error) {
	return http.StatusNotFound, nil, err
}

func OK(body []byte) (int, []byte, error) {
	return http.StatusOK, body, nil
}

func Created(body []byte) (int, []byte, error) {
	return http.StatusCreated, body, nil
}

func NoContent() (int, []byte, error) {
	return http.StatusNoContent, nil, nil
}
