package ecommerce

import (
	"context"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/orders"
)

func (m *Module) CreateStoreOrder(ctx context.Context, storeID string, order *orders.Order) (ID string, err error) {
	//Make sure that the order gts created to the id that its passed, and so you dont
	// have to put the storeID when calling the module
	order.StoreID = storeID

	return m.CreateOrder(ctx, order)
}

func (m *Module) CreateProfileOrder(ctx context.Context, profileID string, order *orders.Order) (ID string, err error) {
	//Make sure that the order gts created to the id that its passed, and so you dont
	// have to put the storeID when calling the module
	order.CustomerID = profileID

	return m.CreateOrder(ctx, order)
}

// CreateOrder creates a buy order on the system, when the order its created a @OrderCreatedEvent will be
// emitted
func (m *Module) CreateOrder(ctx context.Context, order *orders.Order) (ID string, err error) {

	return
}

func (m *Module) GetStoreOrders(ctx context.Context, storeID string) (list []orders.Order, err error) {
	return m.ordersService.GetOrderByStoreID(ctx, storeID)
}

func (m *Module) GetCostumerOrders(ctx context.Context, costumerID string) (list []orders.Order, err error) {
	return m.ordersService.GetOrderByStoreID(ctx, costumerID)
}
