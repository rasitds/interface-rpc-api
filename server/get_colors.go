package server

import (
	"context"
	"fmt"
	api "sencha-twirp-rpc/rpc/colors"

	"github.com/twitchtv/twirp"
)

func (s *ColorsServer) GetColors(ctx context.Context, req *api.GetColorsRequest) (*api.GetColorsResponse, error) {
	if err := validateGetColorsRequest(req); err != nil {
		return nil, err
	}

	colorData, err := s.backend.GetColors(ctx, req.Name)

	if err != nil {
		fmt.Println("Error", err)
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.GetColorsResponse{ColorData: colorData}, nil
}

func validateGetColorsRequest(req *api.GetColorsRequest) error {
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	return nil
}
