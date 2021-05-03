package presentors

import (
	"goarch/app/domain"
	"goarch/app/presentors/html"
	"goarch/app/presentors/json"
	"net/http"
)

type Builder struct{}

func (b *Builder) Make(presenterType string, w http.ResponseWriter) domain.Presenter {
	switch presenterType {
	case "html":
		return &html.Presenter{W: w}
	default:
		return &json.Presenter{W: w}
	}
}

func PresentersFactory() domain.Presenters {
	return &Builder{}
}
