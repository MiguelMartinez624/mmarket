package users

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildUsersModule(client *mongo.Client, r *mux.Router) *Module {
	users := BuildModule()
	return users
}
