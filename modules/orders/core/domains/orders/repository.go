package orders

import "context"

type Repository interface {
	SaveOrder(ctx context.Context, order *Order) (ID string, err error)

	UpdateOrder(ctx context.Context, ID string, order *Order) (ok bool, err error)
}
