package persistency

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/products"
	"go.mongodb.org/mongo-driver/bson"
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
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return false, err
	}
	fmt.Printf("Updating ID %v with %v \n", ID, store)
	filter := bson.M{"_id": id}
	result, err := s.db.UpdateOne(ctx, filter, store)
	if err != nil {
		return false, err
	}
	fmt.Println(result.ModifiedCount)
	success = result.ModifiedCount == 1
	return success, nil
}

func (s *MongoDBProductsRepository) GetAll(ctx context.Context) (products []*products.Product, err error) {
	return
}

func (s *MongoDBProductsRepository) GetAllByStoreID(ctx context.Context, storeID string) (list []*products.Product, err error) {

	cursor, err := s.db.Find(ctx, bson.M{"store_id": storeID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
func (s *MongoDBProductsRepository) GetByID(ctx context.Context, ID string) (item *products.Product, err error) {
	monId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = s.db.FindOne(ctx, bson.M{"id": monId}).Decode(&item)
	if err != nil {
		switch err.Error() {
		case "mongo: no documents in result":
			return nil, err
		default:
			return nil, err
		}

	}

	return item, nil
}
