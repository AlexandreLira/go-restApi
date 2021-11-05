package main

import "github.com/hyperyuri/webapi-with-go/server"

func main() {
	server := server.NewServer()

	server.Run()
}
