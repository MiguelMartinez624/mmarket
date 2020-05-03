package order

import (
	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/modules/orders/core"
	"github.com/miguelmartinez624/mmarket/modules/orders/infrastructure/gateways"
	"github.com/miguelmartinez624/mmarket/modules/orders/infrastructure/persistency"

	"go.mongodb.org/mongo-driver/mongo"
)

func BuildModule(client *mongo.Client, r *mux.Router) *core.Module {

	mongoOrdrsRepo := persistency.NewMongoDBOrdersRepository(client.Database("m_market").Collection("orders"))
	module := core.NewModule(mongoOrdrsRepo)

	//Http Controller
	httpController := gateways.NewHttpController(module)

	//we add endpoints here to mux
	r.HandleFunc("/stores/{store_id}/orders", httpController.CreateOrder).Methods("POST")
	r.HandleFunc("/stores/{store_id}/orders", httpController.GetStoreOrders).Methods("GET")

	return module
}
