package jsonapi

import (
	"encoding/json"
	"goarch/app/domain/models"
)

type item struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (i *item) model() *models.Item {
	return &models.Item{
		Id:          i.Id,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
	}
}

func makeItem(i *models.Item) *item {
	return &item{
		Id:          i.Id,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
	}
}

func MarshalItem(item *models.Item) ([]byte, error) {
	return json.Marshal(item)
}

func MarshalItems(items []*models.Item) ([]byte, error) {
	result := make([]*item, len(items))

	for i, item := range items {
		result[i] = makeItem(item)
	}

	return json.Marshal(result)
}

func UnmarshalItem(body []byte) (*models.Item, error) {
	item := &item{}

	err := json.Unmarshal(body, item)
	if err != nil {
		return nil, err
	}

	return item.model(), nil
}
