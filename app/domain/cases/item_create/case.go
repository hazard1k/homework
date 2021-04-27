package item_create

import (
	"goarch/app/domain"
	"goarch/app/domain/models"
)

func Run(c domain.Connection, item *models.Item) (*models.Item, error) {
	item.Id = ""
	return c.Item().Create(item)
}
