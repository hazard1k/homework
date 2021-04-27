package item_get

import (
	"goarch/app/domain"
	"goarch/app/domain/models"
)

func Run(c domain.Connection, id string) (*models.Item, error) {
	return c.Item().Get(id)
}
