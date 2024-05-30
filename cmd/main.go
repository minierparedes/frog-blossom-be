package main

import (
	"log"

	"github.com/reflection/frog-blossom-be/cmd/api"
)

func main() {
	server := api.NewServer(":8080", nil)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
