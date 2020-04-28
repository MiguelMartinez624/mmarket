package repositories

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/common/models"
)

type Repository interface {
	getAll(ctx context.Context) (list []models.Entity, err error)

	save(ctx context.Context, entity models.Entity) (ID string, err error)

	update(ctx context.Context, ID string, entity models.Entity) (ok bool, err error)

	delete(ctx context.Context, ID string) (ok bool, err error)

	getByID(ctx context.Context, ID string) (entity *models.Entity, err error)
}
