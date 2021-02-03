package main

import (
	"github.com/ilovelili/auth0example/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app.Init()
	Serve()
}
