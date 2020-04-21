package persistency

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/mmarket/modules/authentication/core/domains/accounts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBAccountsRepository struct {
	db *mongo.Collection
}

func NewMongoDBAccountsRepository(db *mongo.Collection) *MongoDBAccountsRepository {
	return &MongoDBAccountsRepository{db: db}
}

func (r *MongoDBAccountsRepository) SaveAccount(ctx context.Context, cre *accounts.Account) (ID string, err error) {
	result, err := r.db.InsertOne(ctx, cre)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return fmt.Sprintf("%v", result.InsertedID), nil

}

func (r *MongoDBAccountsRepository) GetAccountsByUserName(ctx context.Context, username string) (account *accounts.Account, err error) {
	err = r.db.FindOne(ctx, bson.M{"username": username}).Decode(&account)
	if err != nil {
		fmt.Println(err)
		switch err.Error() {
		case "mongo: no documents in result":
			return nil, accounts.AccountDontExist{}
		default:
			return nil, err
		}

	}

	return account, nil
}
func (r *MongoDBAccountsRepository) GetAccountsByValidationHash(ctx context.Context, hash string) (account *accounts.Account, err error) {
	err = r.db.FindOne(ctx, bson.M{"verification_hash": hash}).Decode(&account)
	if err != nil {
		fmt.Println(err)
		switch err.Error() {
		case "mongo: no documents in result":
			return nil, accounts.AccountDontExist{}
		default:
			return nil, err
		}

	}

	return account, nil
}
