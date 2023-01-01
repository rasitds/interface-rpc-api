package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sencha-twirp-rpc/internal/backend/postgresql"
	"sencha-twirp-rpc/internal/db"
	"sencha-twirp-rpc/internal/models"
	api "sencha-twirp-rpc/rpc/themes"
	"sencha-twirp-rpc/server"

	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load()

	if env_err != nil {
		log.Fatalf("Some error occured while loading env file. Err: %s", env_err)
	}

	fmt.Println("dbClient")

	dbClient, err := db.NewPostgreSQLConnect(db.PostgreSQLConfig{Host: os.Getenv("POSTGRES_HOST"), Port: 5432, Username: os.Getenv("POSTGRES_USERNAME"), Password: os.Getenv("POSTGRES_PASSWORD"), DbName: os.Getenv("POSTGRES_DB")})

	fmt.Println("dbClient error handling")

	if err != nil {
		log.Fatal("Error while creating db connection pool. Error:", err)
	}

	fmt.Println("AutoMigrate")

	err = models.AutoMigrate(dbClient)
	if err != nil {
		return
	}

	fmt.Println("AutoMigrate Completed")

	s := server.NewThemesServer(postgresql.NewPostgreSQLBackend(dbClient))
	twirpHandler := api.NewThemesServer(s)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	fmt.Println("PathPrefix", twirpHandler.PathPrefix())

	http.ListenAndServe(":8080", mux)
}
