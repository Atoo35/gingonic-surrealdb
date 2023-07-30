package database

import (
	"log"

	"github.com/surrealdb/surrealdb.go"
)

var DB *surrealdb.DB

func Connect(connString, username, password string) {
	var err error
	DB, err = surrealdb.New(connString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	_, err = DB.Signin(map[string]interface{}{
		"user": username,
		"pass": password,
	})
	if err != nil {
		log.Fatalf("Error signing in: %s", err)
	}

	if _, err = DB.Use("test", "tasks"); err != nil {
		log.Fatalf("Error using database: %s", err)
	}

	log.Printf("Connected to db with namespace %s and collection %s", "test", "tasks")
}
