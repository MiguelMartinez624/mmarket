package persistency

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/stores"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStoresRepository struct {
	db *mongo.Collection
}

func NewMongoDBStoresRepository(db *mongo.Collection) *MongoDBStoresRepository {
	return &MongoDBStoresRepository{db: db}
}
func (s *MongoDBStoresRepository) GetStoresByProfileID(ctx context.Context, profileID string) (list []stores.Store, err error) {

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
	return
}
func (s *MongoDBStoresRepository) Update(ctx context.Context, ID string, store *stores.Store) (success bool, err error) {
	return
}
