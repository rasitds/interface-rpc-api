package server

import (
	"context"
	"fmt"
	api "sencha-twirp-rpc/rpc/themes"
	"sencha-twirp-rpc/server/helper"

	"github.com/twitchtv/twirp"
)

func (s *ThemesServer) GetTheme(ctx context.Context, req *api.GetThemeRequest) (*api.GetThemeResponse, error) {
	if err := validateGetThemeRequest(req); err != nil {
		return nil, err
	}

	themeData, err := s.backend.GetTheme(ctx, req.Name)

	if err != nil {
		fmt.Println("GetTheme Error:", err)
		return nil, twirp.InternalErrorWith(err)
	}

	convertedData := helper.ConvertToThemeModel(themeData)

	return &api.GetThemeResponse{ThemeData: convertedData}, nil
}

func validateGetThemeRequest(req *api.GetThemeRequest) error {
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	return nil
}
