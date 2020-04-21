package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/modules/authentication/infrastructure/controllers"
	"github.com/miguelmartinez624/mmarket/modules/authentication/infrastructure/persistency"
	"github.com/miguelmartinez624/mmarket/modules/authentication/infrastructure/utils"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gompany/core/authentication/facade"
)

const DB_URI = "mongodb://localhost:27017"

func BuildAuthModule(client *mongo.Client, r *mux.Router) *facade.Authentication {

	mongoCredsRepo := persistency.NewMongoDBAccountsRepository(client.Database("gompany").Collection("accounts"))
	bcryptEncripter := utils.BcryptEncripter{}

	auth := facade.NewAuthentication(mongoCredsRepo, bcryptEncripter)

	//Http Controller
	httpController := controllers.NewAuthHTTP(auth)

	//we add endpoints here to mux

	r.HandleFunc("/signin", httpController.Signin).Methods("POST")
	r.HandleFunc("/signup", httpController.SignUp).Methods("POST")

	http.Handle("/", r)
	return auth
}
