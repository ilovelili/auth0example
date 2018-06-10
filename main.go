package main

import (
	"log"

	"github.com/ilovelili/auth0example/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	app.Init()
	Serve()
}
