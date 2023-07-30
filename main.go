package main

import (
	"log"

	"github.com/Atoo35/gingonic-surrealdb/api/routes"
	"github.com/Atoo35/gingonic-surrealdb/database"
)

func init() {
	database.Connect("ws://localhost:8000/rpc", "root", "root")
}

func main() {
	log.Printf("Woohooo")
	router := routes.SetupRoutes(database.DB)

	log.Println("Starting server on port 8080")
	log.Fatal(router.Run(":8080"))
	defer database.DB.Close()
}
