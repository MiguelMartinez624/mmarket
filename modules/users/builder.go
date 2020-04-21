// package users

// import (
// 	"github.com/gorilla/mux"
// 	users "github.com/miguelmartinez624/mmarket/modules/users/core"
// 	"github.com/miguelmartinez624/mmarket/modules/users/infrastructure/persistency"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func BuildUsersModule(client *mongo.Client, r *mux.Router) *Module {

// 	mongoStore := persistency.NewMongoDBProfileStoreRepository(client.Database("m_market"))
// 	users := users.BuildModule(mongoStore)
// 	return users
// }
