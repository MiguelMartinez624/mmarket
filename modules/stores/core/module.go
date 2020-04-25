package core

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/products"
	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/stores"
	"github.com/miguelmartinez624/mmarket/modules/stores/core/externals"
)

type Module struct {
	storeService    *stores.Service
	productsService *products.Service
	proiflesModule  externals.ProfileModule
}

func NewModule(storeRepo stores.Repository, productsRepo products.Repository) *Module {
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
