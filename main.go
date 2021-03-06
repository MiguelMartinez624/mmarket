package main

import (
	"context"
	"fmt"
	"github.com/miguelmartinez624/mmarket/modules/ecommerce"
	"github.com/miguelmartinez624/mmarket/modules/users"
	"github.com/miguelmartinez624/mmarket/nodos"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/miguelmartinez624/mmarket/middlewares"
	auth "github.com/miguelmartinez624/mmarket/modules/authentication"
)

const DB_URI = "mongodb://localhost:27017"

func main() {
	client, cancel := ConnectMongoDB(DB_URI)
	defer cancel()
	r := mux.NewRouter()
	//AuthenticationModule
	authModule, authCell := auth.BuildAuthModule(client, r)
	_, usersCell := users.BuildUsersModule(client, r)
	_, ecommerceCell := ecommerce.BuildModule(client, r)
	//storesModule := ecommerce.BuildModule(client, r)
	//ordersModule := orders.BuildModule(client, r)

	// Mount middldeware dependencies
	middlewares.SetAuthModule(authModule)
	//middlewares.SetStoresModule(storesModule)

	manager := nodos.Manager{
		Nodos: []nodos.Neuron{
			{Name: "authentication", Cell: authCell},
			{Name: "users", Cell: usersCell},
			{Name: "E-Commerce", Cell: ecommerceCell},
		},
	}

	go manager.Start()
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
	fmt.Println("Up and running")
	log.Fatal(srv.ListenAndServe())
}

func ConnectMongoDB(uri string) (client *mongo.Client, cancel context.CancelFunc) {
	log.Println("Connecting to MongoDB...")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, cancel
}
