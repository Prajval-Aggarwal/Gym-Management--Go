package main

import (
	"log"
	"os"

	"gym/server"
	"gym/server/db"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connection := db.InitDB()

	app := server.NewServer(connection)
	server.ConfigureRoutes(app)
	db.Transfer(connection)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}

}
