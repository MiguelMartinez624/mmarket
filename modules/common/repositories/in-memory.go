package repositories

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/common/models"
)

type InMemory struct {
	database []models.Entity
}

func (r InMemory) GetAll(ctx context.Context) (list []models.Entity, err error)          { return }
func (r InMemory) Save(ctx context.Context, entity models.Entity) (ID string, err error) { return }
func (r InMemory) Update(ctx context.Context, ID string, entity models.Entity) (ok bool, err error) {
	return
}
func (r InMemory) Delete(ctx context.Context, ID string) (ok bool, err error)                { return }
func (r InMemory) GetByID(ctx context.Context, ID string) (entity *models.Entity, err error) { return }
func (r InMemory) GetBy(ctx context.Context, query interface{}) (entity *models.Entity, err error) {
	return
}
