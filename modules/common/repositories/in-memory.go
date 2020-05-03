package repositories

import (
	"context"
)

type InMemory struct {
	database []interface{}
}

func (r InMemory) GetAll(ctx context.Context) (list []interface{}, err error) { return }

func (r InMemory) GetAllBy(ctx context.Context, query interface{}, schema interface{}) (list []interface{}, err error) {

	return
}
func (r InMemory) Save(ctx context.Context, entity interface{}) (ID string, err error) { return }
func (r InMemory) Update(ctx context.Context, ID string, entity interface{}) (ok bool, err error) {
	return
}
func (r InMemory) Delete(ctx context.Context, ID string) (ok bool, err error)             { return }
func (r InMemory) GetByID(ctx context.Context, ID string) (entity interface{}, err error) { return }
func (r InMemory) GetBy(ctx context.Context, query interface{}) (entity interface{}, err error) {
	return
}
