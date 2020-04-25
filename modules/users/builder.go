package users

import (
	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/middlewares"
	users "github.com/miguelmartinez624/mmarket/modules/users/core"
	"github.com/miguelmartinez624/mmarket/modules/users/infrastructure/gateway"
	"github.com/miguelmartinez624/mmarket/modules/users/infrastructure/persistency"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildUsersModule(client *mongo.Client, r *mux.Router) *users.Module {

	mongoStore := persistency.NewMongoDBProfileStoreRepository(client.Database("m_market"))
	users := users.BuildModule(mongoStore)

	//Http Controller
	httpController := gateway.NewHttpController(users)

	r.Handle("/users/me", middlewares.IsAuthorized(httpController.Me)).Methods("GET")

	return users
}
