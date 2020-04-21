package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gorilla/mux"
	
	auth "github.com/miguelmartinez624/mmarket/modules/authentication"
	"github.com/miguelmartinez624/mmarket/modules/users"
	"github.com/miguelmartinez624/mmarket/connections"
)

func main() {
	client, cancel := ConnectMongoDB("")
	defer cancel()
	r := mux.NewRouter()
	//AuthenticationModule
	authModule := auth.BuildAuthModule(client, r)

	usersModule := users.BuildUsersModule(client, r)

	authModule.ConnectToProfiles(connections.AuthToProfileConnection(usersModule))
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
