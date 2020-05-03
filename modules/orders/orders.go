package orders

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

	//Store manage orders
	r.HandleFunc("/stores/{store_id}/orders", httpController.CreateStoreOrder).Methods("POST")
	r.HandleFunc("/stores/{store_id}/orders", httpController.GetStoreOrders).Methods("GET")

	//Profile management orders
	r.HandleFunc("/profiles/{profile_id}/orders", httpController.CreateCostumerOrder).Methods("POST")
	r.HandleFunc("/profiles/{profile_id}/orders", httpController.GetCostumerOrders).Methods("GET")

	return module
}
