package html

import (
	"html/template"
	"net/http"
)

type Presenter struct {
	W        http.ResponseWriter
	entity   interface{}
	template *template.Template
}

func (p *Presenter) RenderError(status int, err error) {
	p.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.W.WriteHeader(status)
	tmpl, _ := template.ParseFiles("app/presentors/html/error.tpl")
	tmpl.Execute(p.W, err.Error())
}

func (p *Presenter) RenderSuccess(status int) {
	p.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.W.WriteHeader(status)
	p.template.Execute(p.W, p.entity)
}
