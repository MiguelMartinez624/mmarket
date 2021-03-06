package persistency

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/stores"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStoresRepository struct {
	db *mongo.Collection
}

func NewMongoDBStoresRepository(db *mongo.Collection) *MongoDBStoresRepository {
	return &MongoDBStoresRepository{db: db}
}

func (s *MongoDBStoresRepository) GetStoresByProfileID(ctx context.Context, profileID string) (list []stores.Store, err error) {
	fmt.Println(profileID)
	cursor, err := s.db.Find(ctx, bson.M{"profile_id": profileID})
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

func (s *MongoDBStoresRepository) Save(ctx context.Context, store *stores.Store) (ID string, err error) {
	result, err := s.db.InsertOne(ctx, store)
	if err != nil {
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}

func (s *MongoDBStoresRepository) Update(ctx context.Context, ID string, store *stores.Store) (success bool, err error) {
	return
}

func (s *MongoDBStoresRepository) GetByID(ctx context.Context, ID string) (item *stores.Store, err error) {
	monId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = s.db.FindOne(ctx, bson.M{"_id": monId}).Decode(&item)
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
func (s *MongoDBStoresRepository) GetStoreByIDAndProfileID(ctx context.Context, storeID string, profileID string) (item *stores.Store, err error) {
	monId, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, err
	}

	err = s.db.FindOne(ctx, bson.M{"_id": monId, "profile_id": profileID}).Decode(&item)
	if err != nil {
		switch err.Error() {
		case "mongo: no documents in result":
			return nil, stores.ErrStoreNotFound{}
		default:
			return nil, err
		}

	}

	return item, nil
}
