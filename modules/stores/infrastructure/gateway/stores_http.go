package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/modules/stores/core"
	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/products"
	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/stores"
)

type HttpController struct {
	stores *core.Module
}

func NewHttpController(stores *core.Module) *HttpController {
	return &HttpController{stores: stores}
}

func (a *HttpController) GetUserStores(w http.ResponseWriter, r *http.Request) {
	profileID := mux.Vars(r)["profile_id"]

	storeList, err := a.stores.GetUserStores(r.Context(), profileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&storeList)

}

func (a *HttpController) CreateStore(w http.ResponseWriter, r *http.Request) {
	var store stores.Store

	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(store)
	ID, err := a.stores.CreateStore(r.Context(), &store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "create new store with ID: %v", ID)
}

func (a *HttpController) CreateStoreProduct(w http.ResponseWriter, r *http.Request) {
	var p products.Product

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ID, err := a.stores.CreateStoreProduct(r.Context(), &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "create new store with ID: %v", ID)
}

func (a *HttpController) GetStoreProduct(w http.ResponseWriter, r *http.Request) {
	storeID := mux.Vars(r)["store_id"]

	list, err := a.stores.GetStoreProducts(r.Context(), storeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&list)
}
