package persistency

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/common/repositories"
	"github.com/miguelmartinez624/mmarket/modules/orders/core/domains/orders"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBOrdersRepository struct {
	*repositories.MongoDB
	db *mongo.Collection
}

func NewMongoDBAccountsRepository(db *mongo.Collection) *MongoDBOrdersRepository {
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

func (r *MongoDBOrdersRepository) UpdateOrder(ctx context.Context, ID string, order *orders.Order) (ok bool, err error) {
	return r.Update(ctx, ID, order)
}
