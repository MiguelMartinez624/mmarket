package repositories

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/common/models"
)

type MongoDB struct {
}

func (r MongoDB) GetAll(ctx context.Context) (list []models.Entity, err error)          { return }
func (r MongoDB) Save(ctx context.Context, entity models.Entity) (ID string, err error) { return }
func (r MongoDB) Update(ctx context.Context, ID string, entity models.Entity) (ok bool, err error) {
	return
}
func (r MongoDB) Delete(ctx context.Context, ID string) (ok bool, err error)                { return }
func (r MongoDB) GetByID(ctx context.Context, ID string) (entity *models.Entity, err error) { return }
