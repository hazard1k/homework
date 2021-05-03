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

func (r *itemRepository) GetWhereCategory(category string) ([]*models.Item, error) {
	var items []*item

	ctx := r.getContext()
	cur, err := r.collection().Find(ctx, &bson.M{"category": category})
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

func (r *itemRepository) Store(i *models.Item) (*models.Item, error) {

	properties := makeItemProperties(i)

	if i.Id == "" {
		return nil, fmt.Errorf("item id is empty")
	}

	bsonId, err := primitive.ObjectIDFromHex(i.Id)

	if err != nil {
		return nil, fmt.Errorf("unable to convert item id [%s] to bson id : %s", i.Id, err)
	}

	c := &item{
		Id:         bsonId,
		Properties: properties,
	}

	if _, err := r.collection().ReplaceOne(r.getContext(), &bson.M{"_id": bsonId}, c); err != nil {
		return nil, err
	}

	return i, nil

}

type item struct {
	Id         primitive.ObjectID `bson:"_id"`
	Properties properties         `bson:"properties"`
}

type properties struct {
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Article     string    `bson:"article"`
	Category    string    `bson:"category"`
	Price       itemPrice `bson:"price"`
}

type itemPrice struct {
	Base       float64 `bson:"base"`
	Discounted float64 `bson:"discounted"`
}

func (i *item) model() *models.Item {
	return &models.Item{
		Id:          i.Id.Hex(),
		Name:        i.Properties.Name,
		Article:     i.Properties.Article,
		Category:    i.Properties.Category,
		Description: i.Properties.Description,
		Price: models.ItemPrice{
			Base:       i.Properties.Price.Base,
			Discounted: i.Properties.Price.Discounted,
		},
	}
}

func makeItemProperties(i *models.Item) properties {
	return properties{
		Name:        i.Name,
		Description: i.Description,
		Article:     i.Article,
		Category:    i.Category,
		Price: itemPrice{
			Base:       i.Price.Base,
			Discounted: i.Price.Discounted,
		},
	}
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

	props := makeItemProperties(i)

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
