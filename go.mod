module github.com/miguelmartinez624/mmarket

go 1.13

replace github.com/gompany/core/authentication => ../../gompany/core/authentication

require (
	github.com/gompany/core/authentication v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/miguelmartinez624/web-service-seed v0.0.0-20200419164957-1fd9bd1703f1 // indirect
	go.mongodb.org/mongo-driver v1.3.2
	golang.org/x/crypto v0.0.0-20200414173820-0848c9571904
)
