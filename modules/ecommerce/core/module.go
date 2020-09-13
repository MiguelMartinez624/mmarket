package ecommerce

import (
	"context"
	"fmt"
	"sync"

	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/dto"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/externals"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/products"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/stores"
)

type Module struct {
	storeService    *stores.Service
	productsService *products.Service
	proiflesModule  externals.ProfileModule
}

func Build(storeRepo stores.Repository, productsRepo products.Repository) *Module {
	storeService := stores.NewService(storeRepo)
	productService := products.NewService(productsRepo)

	//Module
	modules := Module{storeService: storeService, productsService: productService}
	return &modules
}

func (m *Module) ConnectToProfiles(connection externals.ProfileModule) {
	m.proiflesModule = connection
}

func (m *Module) CreateStore(ctx context.Context, store *stores.Store) (ID string, err error) {

	_, err = m.proiflesModule.GetProfileByID(ctx, store.ProfileID)
	if err != nil {
		return "", err
	}

	return m.storeService.CreateStore(ctx, store)
}

func (m *Module) GetUserStores(ctx context.Context, profileID string) (list []stores.Store, err error) {

	return m.storeService.GetUserStores(ctx, profileID)
}

func (m *Module) CreateStoreProduct(ctx context.Context, product *products.Product) (ID string, err error) {

	_, err = m.storeService.GetStoreByID(ctx, product.StoreID)
	if err != nil {
		return "", err
	}

	return m.productsService.CreateProduct(ctx, product)
}

func (m *Module) GetStoreProducts(ctx context.Context, storeID string) (list []*products.Product, err error) {
	return m.productsService.GetProductsByStoreID(ctx, storeID)
}

func (m *Module) GetStoreByIDAndProfileID(ctx context.Context, storeID string, profileID string) (store *stores.Store, err error) {
	return m.storeService.GetStoreByIDAndProfileID(ctx, storeID, profileID)
}

func (m *Module) UpdateProduct(ctx context.Context, productID string, product *products.Product) (success bool, err error) {
	return m.productsService.UpdateProduct(ctx, productID, product)
}

func (m *Module) CheckProductsAvailaility(ctx context.Context, consult []dto.ProductStockConsult) (ok bool, err error) {

	var wg sync.WaitGroup
	wg.Add(len(consult))
	invalidCount := 0

	for i := range consult {
		copyC := consult[i]
		go func() {
			wg.Done()
			av, err := m.productsService.CheckAvailability(ctx, copyC.ProductID, copyC.Quantity)
			if err != nil {
				fmt.Println(err)
			}

			if !av {
				invalidCount++
			}
			copyC.Aval = av
		}()
	}

	if invalidCount > 0 {
		return false, nil
	}

	return

}
