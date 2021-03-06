package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/orders"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core"
)

type HttpController struct {
	stores *ecommerce.Module

}

func NewHttpController(stores *ecommerce.Module) *HttpController {
	return &HttpController{stores: stores}
}

func (c *HttpController) CreateStoreOrder(w http.ResponseWriter, r *http.Request) {
	storeID := mux.Vars(r)["store_id"]
	var order orders.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ID, err := c.stores.CreateStoreOrder(r.Context(), storeID, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Created order with ID: %v", ID)
}

func (c *HttpController) CreateCostumerOrder(w http.ResponseWriter, r *http.Request) {
	profileID := mux.Vars(r)["profile_id"]
	var order orders.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ID, err := c.stores.CreateProfileOrder(r.Context(), profileID, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Created order with ID: %v", ID)
}

func (c *HttpController) GetStoreOrders(w http.ResponseWriter, r *http.Request) {
	storeID := mux.Vars(r)["store_id"]

	list, err := c.stores.GetStoreOrders(r.Context(), storeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func (c *HttpController) GetCostumerOrders(w http.ResponseWriter, r *http.Request) {
	profileID := mux.Vars(r)["profile_id"]

	list, err := c.stores.GetCostumerOrders(r.Context(), profileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(list)
}
