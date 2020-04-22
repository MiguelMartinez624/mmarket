package core

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/stores"
	"github.com/miguelmartinez624/mmarket/modules/stores/core/externals"
)

type Module struct {
	storeService   *stores.Service
	proiflesModule externals.ProfileModule
}

func NewModule(storeRepo stores.Repository) *Module {
	storeService := stores.NewService(storeRepo)

	modules := Module{storeService: storeService}
	return &modules
}

func (m *Module) ConnectToProfiles(connection externals.ProfileModule) {
	m.proiflesModule = connection
}

func (m *Module) CreateStore(ctx context.Context, store *stores.Store) (ID string, err error) {

	_, err = m.proiflesModule.GetProfileByID(ctx, ID)
	if err != nil {
		return "", err
	}

	return m.storeService.CreateStore(ctx, store)
}
