package json

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

func (p *Presenter) Item(item *models.Item) error {
	b, err := json.Marshal(makeItem(item))
	if err == nil {
		p.body = b
	}

	return err
}

func (p *Presenter) Items(items []*models.Item) error {
	result := make([]*item, len(items))

	for i, item := range items {
		result[i] = makeItem(item)
	}

	b, err := json.Marshal(result)
	if err == nil {
		p.body = b
	}

	return err
}

func UnmarshalItem(body []byte) (*models.Item, error) {
	item := &item{}

	err := json.Unmarshal(body, item)
	if err != nil {
		return nil, err
	}

	return item.model(), nil
}
