package core

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/mmarket/modules/orders/core/domains/orders"
	"github.com/miguelmartinez624/mmarket/modules/orders/core/externals"
)

type Module struct {
	ordersService *orders.Service
	stores        externals.StoresModule
	Events        chan interface{}
}

func NewModule(orderRepo orders.Repository) *Module {

	orderService := orders.NewService(orderRepo)

	return &Module{ordersService: orderService}
}

// CreateOrder creates a buy order on the system, when the order its created a @OrderCreatedEvent will be
// emitted
func (m *Module) CreateOrder(ctx context.Context, storeID string, order *orders.Order) (ID string, err error) {
	//Make sure that the order gts created to the id that its passed, and so you dont
	// have to put the storeID when calling the module
	order.StoreID = storeID

	// ask availability to the store
	err = m.askProductsAvailability(order)
	if err != nil {
		return "", err
	}

	created, err := m.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return "", err
	}

	fmt.Print(created)
	event := OrderCreatedEvent{
		OrderID: created.ID,
	}

	m.Events <- event

	return
}

func (m *Module) askProductsAvailability(order *orders.Order) error {
	itemList := make([]externals.RequestItem, len(order.Details.Items))
	//Map item details to request item,
	for index, item := range order.Details.Items {
		itemList[index] = externals.RequestItem{
			ProductID: item.ItemID,
			Quantity:  item.Quantity,
		}
	}

	availibles, err := m.stores.AskForAvailability(order.StoreID, itemList)
	if err != nil {
		return err
	}

	invalidItems := make([]string, 2)
	for _, item := range availibles {
		if !item.Availability {
			invalidItems = append(invalidItems, item.ProductID)
		}
	}

	if len(invalidItems) > 0 {
		return ErrUnavailibleItems{Items: invalidItems}
	}

	return nil
}
