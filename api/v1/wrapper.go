package v1

import (
	"github.com/gorilla/mux"
	"goarch/app/domain"
	"net/http"
)

type handleFunc func(method, path string, callable handlerFunc)
type handlerFunc func(v domain.RouteVars, conn domain.Connection, r *http.Request) (int, []byte, error)

func wrap(r *mux.Router, conn domain.Connection) handleFunc {
	internalWrapper := func(callable handlerFunc) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {

			status, response, err := callable(mux.Vars(r), conn, r)

			if err != nil {
				JsonError(w, status, err)
			} else {
				JsonSuccess(w, response)
			}
		}
	}

	return func(method, path string, callable handlerFunc) {
		r.HandleFunc(path, internalWrapper(callable)).Methods(method)
	}
}
