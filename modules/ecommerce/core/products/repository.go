package products

import "context"

type Repository interface {
	Save(ctx context.Context, product *Product) (ID string, err error)
	GetByID(ctx context.Context, ID string) (product *Product, err error)
	GetAllByStoreID(ctx context.Context, storeID string) (product []*Product, err error)
	GetAll(ctx context.Context) (product []*Product, err error)
	Update(ctx context.Context, ID string, product *Product) (success bool, err error)
}
