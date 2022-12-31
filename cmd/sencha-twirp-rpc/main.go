package main

import (
	"fmt"
	"net/http"
	"sencha-twirp-rpc/internal/backend/postgresql"
	api "sencha-twirp-rpc/rpc/colors"
	"sencha-twirp-rpc/server"
)

func main() {
	s := server.NewColorsServer(postgresql.NewPostgreSQLBackend())
	twirpHandler := api.NewColorsServer(s)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	fmt.Println("PathPrefix", twirpHandler.PathPrefix())

	http.ListenAndServe(":8080", mux)
}
