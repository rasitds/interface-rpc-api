package server

import (
	"context"
	"fmt"
	api "sencha-twirp-rpc/rpc/themes"
	"sencha-twirp-rpc/server/helper"

	"github.com/twitchtv/twirp"
)

func (s *ThemesServer) GetThemes(ctx context.Context, req *api.GetThemesRequest) (*api.GetThemesResponse, error) {
	themesData, err := s.backend.GetThemes(ctx)

	if err != nil {
		fmt.Println("GetThemes Error:", err)
		return nil, twirp.InternalErrorWith(err)
	}

	var batch []*api.Theme
	for _, model := range themesData {
		theme := helper.ConvertToThemeModel(model)
		batch = append(batch, theme)
	}

	return &api.GetThemesResponse{ThemeData: batch}, nil
}
