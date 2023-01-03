package server

import (
	"context"
	"fmt"

	api "sencha-twirp-rpc/rpc/themes"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ThemesServer) AddThemes(ctx context.Context, req *api.AddThemesRequest) (*emptypb.Empty, error) {
	err := s.backend.AddThemes(ctx)
	if err != nil {
		fmt.Println("AddThemes Error:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
