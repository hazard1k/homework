package json

import (
	"encoding/json"
	"net/http"
)

const Header = "application/json"

type Presenter struct {
	W    http.ResponseWriter
	body []byte
}

func (p *Presenter) RenderError(status int, err error) {
	p.W.Header().Set("Content-Type", Header)
	p.W.WriteHeader(status)
	e := NewErrorResponse(status, err)
	b, _ := json.Marshal(e)
	p.W.Write(b)
}

func (p *Presenter) RenderSuccess(status int) {
	p.W.Header().Set("Content-Type", Header)
	p.W.WriteHeader(status)
	p.W.Write(p.body)
}
