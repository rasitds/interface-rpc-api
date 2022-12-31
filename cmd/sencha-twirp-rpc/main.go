package main

import (
	"fmt"
	"net/http"
	"sencha-twirp-rpc/internal/backend/postgresql"
	api "sencha-twirp-rpc/rpc/themes"
	"sencha-twirp-rpc/server"
)

func main() {
	s := server.NewThemesServer(postgresql.NewPostgreSQLBackend())
	twirpHandler := api.NewThemesServer(s)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	fmt.Println("PathPrefix", twirpHandler.PathPrefix())

	http.ListenAndServe(":8080", mux)
}
