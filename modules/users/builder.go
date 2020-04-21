package users

import (
	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/modules/users/infrastructure/persistency"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildUsersModule(client *mongo.Client, r *mux.Router) *Module {

	mongoStore := persistency.NewMongoDBProfileStoreRepository(client.Database("m_market"))
	users := BuildModule(mongoStore)
	return users
}
