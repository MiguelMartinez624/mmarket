package stores

import (
	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/middlewares"
	"github.com/miguelmartinez624/mmarket/modules/stores/core"
	"github.com/miguelmartinez624/mmarket/modules/stores/infrastructure/gateway"
	"github.com/miguelmartinez624/mmarket/modules/stores/infrastructure/persistency"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildModule(client *mongo.Client, r *mux.Router) *core.Module {

	mongoStoresRepo := persistency.NewMongoDBStoresRepository(client.Database("m_market").Collection("stores"))
	mongoProductsRepo := persistency.NewMongoDBProductsRepository(client.Database("m_market").Collection("products"))
	module := core.NewModule(mongoStoresRepo, mongoProductsRepo)

	//Http Controller
	httpController := gateway.NewHttpController(module)

	//we add endpoints here to mux
	r.HandleFunc("/{profile_id}/stores", httpController.GetUserStores).Methods("GET") // probably later on make a me stores
	r.HandleFunc("/{profile_id}/stores", middlewares.IsAuthorized(middlewares.OwnResource(httpController.CreateStore))).Methods("POST")

	//Store products endpoints
	r.HandleFunc("/{profile_id}/stores/{store_id}/products", middlewares.IsAuthorized(middlewares.OwnStore(httpController.GetStoreProduct))).Methods("GET")
	r.HandleFunc("/{profile_id}/stores/{store_id}/products", middlewares.IsAuthorized(middlewares.OwnStore(httpController.CreateStoreProduct))).Methods("POST")
	r.HandleFunc("/{profile_id}/stores/{store_id}/products/{product_id}", middlewares.IsAuthorized(middlewares.OwnStore(httpController.UpdateProduct))).Methods("POST")

	return module
}
