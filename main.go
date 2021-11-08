package main

import (
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
