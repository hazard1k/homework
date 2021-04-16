package v1

import (
	"github.com/gorilla/mux"
	"goarch/app/domain"
	"net/http"
)

type handleFunc func(method, path string, callable handlerFunc)
type handlerFunc func(v map[string]string, conn domain.Connection, w http.ResponseWriter, r *http.Request) (int, []byte, error)

func wrap(r *mux.Router, conn domain.Connection) handleFunc {
	internalWrapper := func(callable handlerFunc) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			status, response, err := callable(vars, conn, w, r)

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
