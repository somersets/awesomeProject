package main

import (
	"awesomeProject/infrastructure"
	"log"
)

func main() {
	config := infrastructure.NewConfig()
	api := infrastructure.NewApi(config)

	if err := api.Start(); err != nil {
		log.Fatalf("Failed at starting server...")
	}

}
