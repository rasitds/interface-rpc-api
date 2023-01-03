package helper

import (
	"sencha-twirp-rpc/internal/models"
	api "sencha-twirp-rpc/rpc/themes"
)

func ConvertToThemeModel(t *models.Theme) *api.Theme {
	return &api.Theme{Name: t.Name, Background: t.Background, Foreground: t.Foreground}
}
