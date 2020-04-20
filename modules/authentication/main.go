package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gompany/core/authentication/facade"
	"github.com/gompany/core/authentication/infrastructure/controllers"
	"github.com/gompany/core/authentication/infrastructure/persistency"
	"github.com/gompany/core/authentication/infrastructure/utils"
)

const DB_URI = "mongodb://localhost:27017"

func main() {
	fmt.Println("Golang and Docker")
	client, cancel := ConnectMongoDB(DB_URI)
	defer cancel()

	mongoCredsRepo := persistency.NewMongoDBAccountsRepository(client.Database("gompany").Collection("accounts"))
	bcryptEncripter := utils.BcryptEncripter{}

	auth := facade.NewAuthentication(mongoCredsRepo, bcryptEncripter)

	//Http Controller
	httpController := controllers.NewAuthHTTP(auth)

	//we add endpoints here to mux
	r := mux.NewRouter()
	r.HandleFunc("/signin", httpController.Signin).Methods("POST")
	r.HandleFunc("/signup", httpController.SignUp).Methods("POST")

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Up and running")
	log.Fatal(srv.ListenAndServe())

}

func ConnectMongoDB(uri string) (client *mongo.Client, cancel context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, cancel
}
