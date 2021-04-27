package items_get

import (
	"goarch/app/domain"
	"goarch/app/domain/models"
)

func Run(c domain.Connection) ([]*models.Item, error) {
	return c.Item().GetAll()
}
