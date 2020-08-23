package persistency

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/orders/core/domains/orders"
	"github.com/miguelmartinez624/mmarket/nodos/repositories"
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
	raw, err := r.GetAllBy(ctx, query, orders.Order{})
	if err != nil {
		return nil, err
	}

	// Map to objct this should done inside mongodb TODO MOVE
	list = make([]orders.Order, len(raw))
	for i, o := range raw {
		list[i] = o.(orders.Order)
	}

	return
}

func (r *MongoDBOrdersRepository) GetOrdersByCostumerID(ctx context.Context, costumerID string) (list []orders.Order, err error) {

	query := bson.M{"costumer_id": costumerID}
	raw, err := r.GetAllBy(ctx, query, orders.Order{})
	if err != nil {
		return nil, err
	}

	// Map to objct this should done inside mongodb TODO MOVE
	list = make([]orders.Order, len(raw))
	for i, o := range raw {
		list[i] = o.(orders.Order)
	}

	return
}

func (r *MongoDBOrdersRepository) UpdateOrder(ctx context.Context, ID string, order *orders.Order) (ok bool, err error) {
	return r.Update(ctx, ID, order)
}
