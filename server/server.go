package server

import "sencha-twirp-rpc/internal/backend"

type ColorsServer struct {
	backend backend.Backender
}

func NewColorsServer(backend backend.Backender) *ColorsServer {
	return &ColorsServer{backend: backend}
}
