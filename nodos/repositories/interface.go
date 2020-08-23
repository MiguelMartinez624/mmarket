package repositories

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) (list []interface{}, err error)

	GetAllBy(ctx context.Context, query interface{}, schema interface{}) (list []interface{}, err error)

	Save(ctx context.Context, entity interface{}) (ID string, err error)

	Update(ctx context.Context, ID string, entity interface{}) (ok bool, err error)

	Delete(ctx context.Context, ID string) (ok bool, err error)

	GetByID(ctx context.Context, ID string) (entity *interface{}, err error)

	GetBy(ctx context.Context, query interface{}, output interface{}) (err error)
}
