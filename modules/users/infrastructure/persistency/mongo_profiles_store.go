package persistency

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/mmarket/modules/users/core/domains/profiles"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBProfileStore struct {
	db *mongo.Database
}

func NewMongoDBProfileStoreRepository(db *mongo.Database) *MongoDBProfileStore {
	return &MongoDBProfileStore{db: db}
}
func (s *MongoDBProfileStore) StoreProfile(ctx context.Context, profile *profiles.Profile) (ID string, err error) {

	result, err := s.db.Collection("profiles").InsertOne(ctx, profile)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}

func (s *MongoDBProfileStore) FindProfileByID(ctx context.Context, ID string) (profile *profiles.Profile, err error) {
	err = s.db.Collection("profiles").FindOne(ctx, bson.M{"id": ID}).Decode(&profile)
	if err != nil {
		fmt.Println(err)
		switch err.Error() {
		case "mongo: no documents in result":
			return nil, profiles.ProfileDontFoundError{}
		default:
			return nil, err
		}

	}

	return
}

func (s *MongoDBProfileStore) FindProfileByAccountID(ctx context.Context, accountId string) (profile *profiles.Profile, err error) {
	err = s.db.Collection("profiles").FindOne(ctx, bson.M{"account_id": accountId}).Decode(&profile)
	if err != nil {
		fmt.Println(err)
		switch err.Error() {
		case "mongo: no documents in result":
			return nil, profiles.ProfileDontFoundError{}
		default:
			return nil, err
		}

	}

	return
}

func (s *MongoDBProfileStore) FindContactByID(ctx context.Context, contactId string) (profile *profiles.ContactInfo, err error) {
	return
}
