package main

import (
	"go-api/cmd/api"
	"log"
)

func main() {
	server := api.NewServer(":3333", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
