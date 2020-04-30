package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/mmarket/modules/authentication/infrastructure/controllers"
	"github.com/miguelmartinez624/mmarket/modules/authentication/infrastructure/persistency"
	"github.com/miguelmartinez624/mmarket/modules/authentication/infrastructure/utils"
	"go.mongodb.org/mongo-driver/mongo"

	authModule "github.com/miguelmartinez624/mmarket/modules/authentication/core"
)

const DB_URI = "mongodb://localhost:27017"

func BuildAuthModule(client *mongo.Client, r *mux.Router) *authModule.Module {

	mongoCredsRepo := persistency.NewMongoDBAccountsRepository(client.Database("m_market").Collection("accounts"))
	bcryptEncripter := utils.BcryptEncripter{}
	jwtToken := &utils.JWTTokenManager{}
	auth := authModule.NewAuthentication(mongoCredsRepo, bcryptEncripter, jwtToken)

	//Http Controller
	httpController := controllers.NewAuthHTTP(auth)

	//we add endpoints here to mux

	r.HandleFunc("/signin", httpController.Signin).Methods("POST", "OPTIONs")
	r.HandleFunc("/signup", httpController.SignUp).Methods("POST", "OPTIONs")
	r.HandleFunc("/validate/{validation_code}", httpController.ValidateAccount).Methods("get")

	http.Handle("/", r)
	return auth
}
