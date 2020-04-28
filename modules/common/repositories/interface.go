package repositories

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/common/models"
)

type Repository interface {
	GetAll(ctx context.Context) (list []models.Entity, err error)

	Save(ctx context.Context, entity models.Entity) (ID string, err error)

	Update(ctx context.Context, ID string, entity models.Entity) (ok bool, err error)

	Delete(ctx context.Context, ID string) (ok bool, err error)

	GetByID(ctx context.Context, ID string) (entity *models.Entity, err error)

	GetBy(ctx context.Context, query interface{}, output interface{}) (err error)
}
