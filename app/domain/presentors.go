package domain

import (
	"goarch/app/domain/models"
	"net/http"
)

type Presenter interface {
	RenderError(status int, err error)
	RenderSuccess(status int)

	Item(item *models.Item) error
	Items(items []*models.Item) error
}

type Presenters interface {
	Make(presenterType string, w http.ResponseWriter) Presenter
}
