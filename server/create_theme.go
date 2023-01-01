package server

import (
	"context"
	api "sencha-twirp-rpc/rpc/themes"
	"sencha-twirp-rpc/server/helper"

	"github.com/twitchtv/twirp"
)

func (s *ThemesServer) CreateTheme(ctx context.Context, req *api.CreateThemeRequest) (*api.GetThemeResponse, error) {
	if err := validateCreateThemeRequest(req); err != nil {
		return nil, err
	}

	theme, err := s.backend.CreateTheme(ctx, req.Name, req.Background, req.Foreground)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	convertedData := helper.ConvertToThemeModel(theme)

	return &api.GetThemeResponse{ThemeData: convertedData}, nil
}

func validateCreateThemeRequest(req *api.CreateThemeRequest) error {
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	if req.Background == "" {
		return twirp.RequiredArgumentError("background")
	}
	if req.Foreground == "" {
		return twirp.RequiredArgumentError("foreground")
	}
	return nil
}
