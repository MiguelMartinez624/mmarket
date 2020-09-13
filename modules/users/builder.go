package users

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/middlewares"
	users "github.com/miguelmartinez624/mmarket/modules/users/core"
	"github.com/miguelmartinez624/mmarket/modules/users/infrastructure/gateway"
	"github.com/miguelmartinez624/mmarket/modules/users/infrastructure/persistency"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildUsersModule(client *mongo.Client, r *mux.Router) (*users.Module, *UsersCell) {

	mongoStore := persistency.NewMongoDBProfileStoreRepository(client.Database("m_market"))
	users := users.BuildModule(mongoStore)
	cell := &UsersCell{
		module: users,
	}
	//Http Controller
	httpController := gateway.NewHttpController(users)

	r.Handle("/users/me", middlewares.IsAuthorized(httpController.Me)).Methods("GET")

	fmt.Println("Users/Profile module running")
	return users, cell
}
