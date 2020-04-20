module github.com/miguelmartinez624/mmarket

go 1.13

replace github.com/gompany/core/authentication => ../../gompany/core/authentication

require (
	github.com/gompany/core/authentication v1.0.0
	github.com/gorilla/mux v1.7.4
	go.mongodb.org/mongo-driver v1.3.2
	golang.org/x/crypto v0.0.0-20200414173820-0848c9571904
)
