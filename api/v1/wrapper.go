package v1

import (
	"github.com/gorilla/mux"
	"goarch/app/domain"
	"net/http"
)

type handleFunc func(method, path string, callable handlerFunc)
type handlerFunc func(v domain.RouteVars, ctx domain.Context, r *http.Request, p domain.Presenter) (int, error)

func wrap(r *mux.Router, ctx domain.Context) handleFunc {
	internalWrapper := func(callable handlerFunc) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {

			p := ctx.PresentersFactory().Make(r.URL.Query().Get("format"), w)

			status, err := callable(mux.Vars(r), ctx, r, p)

			if err != nil {
				p.RenderError(status, err)
			} else {
				p.RenderSuccess(status)
			}
		}
	}

	return func(method, path string, callable handlerFunc) {
		r.HandleFunc(path, internalWrapper(callable)).Methods(method)
	}
}
