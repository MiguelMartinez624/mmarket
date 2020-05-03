package persistency

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/common/repositories"
	"github.com/miguelmartinez624/mmarket/modules/orders/core/domains/orders"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBOrdersRepository struct {
	*repositories.MongoDB
	db *mongo.Collection
}

func NewMongoDBOrdersRepository(db *mongo.Collection) *MongoDBOrdersRepository {
	return &MongoDBOrdersRepository{
		db:      db,
		MongoDB: repositories.NewMongoDBRepo(db),
	}
}

func (r *MongoDBOrdersRepository) SaveOrder(ctx context.Context, order *orders.Order) (ID string, err error) {
	return r.Save(ctx, order)
}

func (r *MongoDBOrdersRepository) GetOrderByID(ctx context.Context, ID string) (order *orders.Order, err error) {
	order = &orders.Order{}
	err = r.GetByID(ctx, ID, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
func (r *MongoDBOrdersRepository) GetOrdersByStoreID(ctx context.Context, storeID string) (list []orders.Order, err error) {

	query := bson.M{"store_id": storeID}
	r.GetAllBy(ctx, query)
	return
}

func (r *MongoDBOrdersRepository) UpdateOrder(ctx context.Context, ID string, order *orders.Order) (ok bool, err error) {
	return r.Update(ctx, ID, order)
}
