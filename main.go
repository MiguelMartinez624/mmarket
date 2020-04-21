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

	"github.com/miguelmartinez624/mmarket/connections"
	auth "github.com/miguelmartinez624/mmarket/modules/authentication"
	"github.com/miguelmartinez624/mmarket/modules/users"
)

const DB_URI = "mongodb://localhost:27017"

func main() {
	client, cancel := ConnectMongoDB(DB_URI)
	defer cancel()
	r := mux.NewRouter()
	//AuthenticationModule
	authModule := auth.BuildAuthModule(client, r)

	usersModule := users.BuildUsersModule(client, r)

	authModule.ConnectToProfiles(connections.AuthToProfileConnection(usersModule))

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
