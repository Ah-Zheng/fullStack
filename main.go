package main

import (
	"log"

	"full-stack/api"
)

func main() {
	server := api.SetupApi()

	if err := server.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
