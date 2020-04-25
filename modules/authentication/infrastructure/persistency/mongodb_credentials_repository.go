package persistency

import (
	"context"
	"fmt"
	"log"

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
	err = r.db.FindOne(ctx, bson.M{"validation_hash": hash}).Decode(&account)
	if err != nil {
		switch err.Error() {
		case "mongo: no documents in result":
			return nil, accounts.AccountDontExist{}
		default:
			return nil, err
		}

	}

	return account, nil
}

func (r *MongoDBAccountsRepository) UpdateAccount(ctx context.Context, ID string, account *accounts.Account) (success bool, err error) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updating ID %v with %v \n", ID, account)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": account.Status}}
	result, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	fmt.Println(result.ModifiedCount)
	success = result.ModifiedCount == 1
	return success, nil
}
