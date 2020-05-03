package gateways

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/modules/orders/core"
	"github.com/miguelmartinez624/mmarket/modules/orders/core/domains/orders"
)

type HttpController struct {
	stores *core.Module
}

func NewHttpController(stores *core.Module) *HttpController {
	return &HttpController{stores: stores}
}

func (c *HttpController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	storeID := mux.Vars(r)["store_id"]
	var order orders.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ID, err := c.stores.CreateOrder(r.Context(), storeID, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Created order with ID: %v", ID)

}
