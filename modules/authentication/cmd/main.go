package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	auth "github.com/miguelmartinez624/mmarket/modules/authentication"
	"github.com/miguelmartinez624/mmarket/modules/nodos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

const DB_URI = "mongodb://localhost:27017"

func main() {
	client, cancel := ConnectMongoDB(DB_URI)
	defer cancel()
	r := mux.NewRouter()

	module := auth.BuildAuthModule(client, r)
	module.SetNotificationHandler(func(ev nodos.Event) {
		log.Println(ev)
	})

	// Service start
	handler := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r)
	// corsObj := handlers.AllowedOrigins([]string{"*"})
	srv := &http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Up and running")
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
