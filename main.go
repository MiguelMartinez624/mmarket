package main

import (
	"context"
	"fmt"
	"log"
	"time"

	auth "github.com/miguelmartinez624/mmarket/modules/authentication"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, cancel := ConnectMongoDB("")
	defer cancel()

	//AuthenticationModule
	authModule := auth.BuildAuthModule(client)
	fmt.Println(authModule)
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
