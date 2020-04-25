package persistency

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBProductsRepository struct {
	db *mongo.Collection
}

func NewMongoDBProductsRepository(db *mongo.Collection) *MongoDBProductsRepository {
	return &MongoDBProductsRepository{db: db}
}

func (s *MongoDBProductsRepository) Save(ctx context.Context, store *products.Product) (ID string, err error) {
	result, err := s.db.InsertOne(ctx, store)
	if err != nil {
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}

func (s *MongoDBProductsRepository) Update(ctx context.Context, ID string, store *products.Product) (success bool, err error) {
	return
}
func (s *MongoDBProductsRepository) GetAll(ctx context.Context) (products []*products.Product, err error) {
	return
}

func (s *MongoDBProductsRepository) GetAllByStoreID(ctx context.Context, storeID string) (products []*products.Product, err error) {
	return
}
func (s *MongoDBProductsRepository) GetByID(ctx context.Context, ID string) (products *products.Product, err error) {
	return
}
