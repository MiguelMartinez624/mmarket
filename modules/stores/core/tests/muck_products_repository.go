package tests

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/products"
)

var DATABASEProducts map[string]*products.Product = make(map[string]*products.Product)

type MuckProductRepository struct{}

func (s *MuckProductRepository) GetByID(ctx context.Context, ID string) (product *products.Product, err error) {

	return
}
func (s *MuckProductRepository) GetAll(ctx context.Context) (list []*products.Product, err error) {

	return list, nil
}
func (s *MuckProductRepository) GetAllByStoreID(ctx context.Context, storeID string) (list []*products.Product, err error) {

	return list, nil
}

func (s *MuckProductRepository) Save(ctx context.Context, p *products.Product) (ID string, err error) {
	p.ID = string(len(DATABASE))
	DATABASEProducts[p.ID] = p
	return p.ID, nil
}
func (s *MuckProductRepository) Update(ctx context.Context, ID string, p *products.Product) (success bool, err error) {
	return true, nil
}
