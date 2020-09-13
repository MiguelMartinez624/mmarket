package gateway

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/products"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/stores"
)

type HttpController struct {
	stores *ecommerce.Module
}

func NewHttpController(stores *ecommerce.Module) *HttpController {
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
	log.Println("Attemptin to create product.")
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

func (a *HttpController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var p products.Product
	pID := mux.Vars(r)["product_id"]
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := a.stores.UpdateProduct(r.Context(), pID, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "product updated: %v", ok)
}
