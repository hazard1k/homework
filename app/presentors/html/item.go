package html

import (
	"goarch/app/domain/models"
	"html/template"
)

func (p *Presenter) Item(item *models.Item) error {
	tmpl, err := template.ParseFiles("app/presentors/html/item.tpl")
	if err == nil {
		p.template = tmpl
		p.entity = item
	}

	return err
}

func (p *Presenter) Items(items []*models.Item) error {
	tmpl, err := template.ParseFiles("app/presentors/html/items.tpl")
	if err == nil {
		p.template = tmpl
		p.entity = items
	}

	return err
}
