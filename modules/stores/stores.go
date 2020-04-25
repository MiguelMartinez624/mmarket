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
	module := core.NewModule(mongoStoresRepo)

	//Http Controller
	httpController := gateway.NewHttpController(module)

	//we add endpoints here to mux
	r.HandleFunc("/{profile_id}/stores", httpController.GetUserStores).Methods("GET") // probably later on make a me stores
	r.HandleFunc("/{profile_id}/stores", middlewares.IsAuthorized(middlewares.OwnResource(httpController.CreateStore))).Methods("POST")

	return module
}
