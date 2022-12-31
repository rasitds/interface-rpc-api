package server

import "sencha-twirp-rpc/internal/backend"

type ThemesServer struct {
	backend backend.Backender
}

func NewThemesServer(backend backend.Backender) *ThemesServer {
	return &ThemesServer{backend: backend}
}
