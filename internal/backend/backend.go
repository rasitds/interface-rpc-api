package backend

import (
	"context"
	"sencha-twirp-rpc/internal/models"
)

type Backender interface {
	GetTheme(ctx context.Context, name string) (*models.Theme, error)
	CreateTheme(ctx context.Context, name string, background string, foreground string) (*models.Theme, error)
}
