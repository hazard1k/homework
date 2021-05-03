package domain

import "goarch/app/domain/repositories"

type Connection interface {
	ItemRepository() repositories.ItemRepository
}
