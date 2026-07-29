# P2P-Messenger
# Description
This messenger is designed with P2P technology. There are also authorization and registration with JWT included. 
# Stack
* Golang
* PostgreSQL
* Migrations
* JWT
* Docker
# Start the app
To start the app you need to download all dependencies with writing `go mod tidy` in your console. After that you need to create .env file in the root and fill all the lines as in the .env.example. After that you can start the app with `go run cmd/app/main.go`