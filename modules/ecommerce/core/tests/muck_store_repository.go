package tests

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/stores"
)

var DATABASE map[string]*stores.Store = make(map[string]*stores.Store)

type MuckStoreRepository struct{}

func (s *MuckStoreRepository) GetStoresByProfileID(ctx context.Context, profilID string) (list []stores.Store, err error) {
	list = make([]stores.Store, 1)
	for _, p := range DATABASE {
		if p.ProfileID == profilID {
			list = append(list, *p)
		}
	}

	return list, nil
}

func (s *MuckStoreRepository) GetAccountByID(ctx context.Context, ID string) (store *stores.Store, err error) {
	for _, p := range DATABASE {
		if p.ID == ID {
			store = p
		}
	}
	if store == nil {
		return nil, nil
	}

	return
}
func (s *MuckStoreRepository) Save(ctx context.Context, store *stores.Store) (ID string, err error) {
	store.ID = string(len(DATABASE))
	DATABASE[store.ID] = store
	return store.ID, nil
}
func (s *MuckStoreRepository) Update(ctx context.Context, ID string, store *stores.Store) (success bool, err error) {
	return true, nil
}

func (s *MuckStoreRepository) GetByID(ctx context.Context, ID string) (item *stores.Store, err error) {

	return
}

func (s *MuckStoreRepository) GetStoreByIDAndProfileID(ctx context.Context, storeID string, profileID string) (item *stores.Store, err error) {

	return
}
