package v1

import (
	"github.com/gorilla/mux"
	"goarch/app/domain"
	"net/http"
)

const ApiPrefix = "/v1"

func Register(r *mux.Router, c domain.Context) {
	router := r.PathPrefix(ApiPrefix).Subrouter()
	handle := wrap(router, c)

	handle(http.MethodGet, "/items", itemsGetAll)
	handle(http.MethodGet, "/items/{id}", itemGet)
	handle(http.MethodPost, "/items", itemCreate)
	handle(http.MethodDelete, "/items/{id}", itemDelete)
}
