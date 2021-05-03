package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goarch/app/domain"
	"goarch/app/domain/repositories"
)

type connection struct {
	client       *mongo.Client
	repositories repos
}

type repos struct {
	itemRepository repositories.ItemRepository
}

func (c *connection) ItemRepository() repositories.ItemRepository {
	return c.repositories.itemRepository
}

func NewConnection(connectionString string) (domain.Connection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	db := client.Database("shop")

	return &connection{
		client: client,
		repositories: repos{
			itemRepository: makeItemRepository(db),
		},
	}, nil
}
