package orders

import "context"

type Repository interface {
	SaveOrder(ctx context.Context, order *Order) (ID string, err error)

	GetOrderByID(ctx context.Context, ID string) (order *Order, err error)

	GetOrdersByStoreID(ctx context.Context, ID string) (list []Order, err error)

	UpdateOrder(ctx context.Context, ID string, order *Order) (ok bool, err error)
}
