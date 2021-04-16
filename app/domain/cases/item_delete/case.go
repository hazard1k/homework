package item_delete

import (
	"goarch/app/domain"
)

func Run(c domain.Connection, id string) error {
	return c.Item().Delete(id)
}
