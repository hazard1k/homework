package v1

import (
	"github.com/gorilla/mux"
	"goarch/app/domain"
	"net/http"
)

const ApiPrefix = "/v1"

func Register(r *mux.Router, connection domain.Connection) {
	router := r.PathPrefix(ApiPrefix).Subrouter()
	handle := wrap(router, connection)

	handle(http.MethodGet, "/items", itemsGetAll)
	handle(http.MethodGet, "/items/{id}", itemGet)
	handle(http.MethodPost, "/items", itemCreate)
	handle(http.MethodDelete, "/items/{id}", itemDelete)
}
