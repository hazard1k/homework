package json

import (
	"net/http"
)

const Header = "application/json"

type Presenter struct {
	W    http.ResponseWriter
	body []byte
}

func (p *Presenter) RenderError(status int, err error) {
	panic("implement me")
}

func (p *Presenter) RenderSuccess(status int) {
	p.W.Header().Set("Content-Type", Header)
	p.W.WriteHeader(status)
	p.W.Write(p.body)
}
