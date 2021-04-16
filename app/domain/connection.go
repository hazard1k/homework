package domain

import "goarch/app/domain/repositories"

type Connection interface {
	Item() repositories.ItemRepository
}
