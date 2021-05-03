package repositories

import (
	"goarch/app/domain/models"
)

type ItemRepository interface {
	GetAll() ([]*models.Item, error)
	GetWhereCategory(category string) ([]*models.Item, error)
	Create(item *models.Item) (*models.Item, error)
	Get(id string) (*models.Item, error)
	Delete(id string) error
	Store(item *models.Item) (*models.Item, error)
}
