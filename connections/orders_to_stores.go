package connections

import (
	"github.com/miguelmartinez624/mmarket/modules/orders/core/externals"
	stores "github.com/miguelmartinez624/mmarket/modules/stores/core"
)

type OrdersToStores struct {
	m *stores.Module
}

func OrdersToStoresConnection(u *stores.Module) *OrdersToStores {
	return &OrdersToStores{m: u}
}

func (c OrdersToStores) AskForAvailability(storeID string, car []externals.RequestItem) (availability []externals.AvailavilityStatus, err error) {
	return
}
