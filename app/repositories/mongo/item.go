package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"goarch/app/domain/models"
	"time"
)

type itemRepository struct {
	db *mongo.Database
}

type item struct {
	Id         primitive.ObjectID `bson:"_id"`
	Properties properties         `bson:"properties"`
}

type properties struct {
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
}

func (i *item) model() *models.Item {
	return &models.Item{
		Id:          i.Id.Hex(),
		Name:        i.Properties.Name,
		Description: i.Properties.Description,
		Price:       i.Properties.Price,
	}
}

func makeItemProperties(i *models.Item) (properties, error) {
	return properties{
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
	}, nil
}

func (r *itemRepository) GetAll() ([]*models.Item, error) {
	var items []*item

	ctx := r.getContext()
	cur, err := r.collection().Find(ctx, &bson.M{})
	if err != nil {
		return nil, fmt.Errorf("unable to find items: %s", err)
	}

	if err := cur.All(ctx, &items); err != nil {
		return nil, fmt.Errorf("unable to fetch items: %s", err)
	}

	modelItems := make([]*models.Item, len(items))

	for idx, i := range items {
		modelItems[idx] = i.model()
	}

	return modelItems, nil
}

func (r *itemRepository) Create(i *models.Item) (*models.Item, error) {

	props, err := makeItemProperties(i)
	if err != nil {
		return nil, err
	}

	insert := struct {
		Properties properties `bson:"properties"`
	}{props}

	result, err := r.collection().InsertOne(r.getContext(), insert)
	if err != nil {
		return nil, fmt.Errorf("unable to insert item: %s", err)
	}

	response := &item{
		Id:         result.InsertedID.(primitive.ObjectID),
		Properties: props,
	}

	return response.model(), nil
}

func (r *itemRepository) Get(id string) (*models.Item, error) {

	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("unable to convert item id to object id")
	}

	item := &item{}

	res := r.collection().FindOne(r.getContext(), &bson.M{"_id": itemId})

	err = res.Decode(item)
	if err != nil {
		return nil, fmt.Errorf("unable to decode item: %s", err)
	}

	return item.model(), nil
}

func (r *itemRepository) Delete(id string) error {

	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("unable to conver id to bson id: %s", err)
	}

	_, err = r.collection().DeleteOne(r.getContext(), &bson.M{"_id": bsonId})
	return err
}

func (r *itemRepository) collection() *mongo.Collection {
	return r.db.Collection("items")
}

func (r *itemRepository) getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	return ctx
}

func makeItemRepository(db *mongo.Database) *itemRepository {
	return &itemRepository{db: db}
}
