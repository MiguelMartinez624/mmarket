package ecommerce

import (
	"fmt"
	"github.com/gorilla/mux"
	ecommerce "github.com/miguelmartinez624/mmarket/modules/ecommerce/core"

	"github.com/miguelmartinez624/mmarket/middlewares"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/infrastructure/gateway"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce/infrastructure/persistency"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildModule(client *mongo.Client, r *mux.Router) (*ecommerce.Module, *EcommerceCell) {

	mongoStoresRepo := persistency.NewMongoDBStoresRepository(client.Database("m_market").Collection("ecommerce"))
	mongoProductsRepo := persistency.NewMongoDBProductsRepository(client.Database("m_market").Collection("products"))
	module := ecommerce.Build(mongoStoresRepo, mongoProductsRepo)

	//Http Controller
	httpController := gateway.NewHttpController(module)
	cell := &EcommerceCell{
		module: module,
	}
	//we add endpoints here to mux
	r.HandleFunc("/{profile_id}/ecommerce", httpController.GetUserStores).Methods("GET") // probably later on make a me ecommerce
	r.HandleFunc("/{profile_id}/ecommerce", middlewares.IsAuthorized(middlewares.OwnResource(httpController.CreateStore))).Methods("POST")

	//Store products endpoints
	r.HandleFunc("/{profile_id}/ecommerce/{store_id}/products", middlewares.IsAuthorized(middlewares.OwnStore(httpController.GetStoreProduct))).Methods("GET")
	r.HandleFunc("/{profile_id}/ecommerce/{store_id}/products", middlewares.IsAuthorized(middlewares.OwnStore(httpController.CreateStoreProduct))).Methods("POST")
	r.HandleFunc("/{profile_id}/ecommerce/{store_id}/products/{product_id}", middlewares.IsAuthorized(middlewares.OwnStore(httpController.UpdateProduct))).Methods("POST")

	fmt.Println("Stores module running ....")
	return module, cell
}
