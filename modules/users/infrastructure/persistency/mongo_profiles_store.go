package persistency

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/mmarket/modules/users/domains/profile"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBProfileStore struct {
	db *mongo.Database
}

func NewMongoDBProfileStoreRepository(db *mongo.Database) *MongoDBProfileStore {
	return &MongoDBProfileStore{db: db}
}
func (s *MongoDBProfileStore) StoreProfile(ctx context.Context, profile *profile.Profile) (ID string, err error) {

	result, err := s.db.Collection("profiles").InsertOne(ctx, profile)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}

func (s *MongoDBProfileStore) FindProfileByID(ctx context.Context, ID string) (profile *profile.Profile, err error) {
	return
}
