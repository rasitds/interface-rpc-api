package backend

import (
	"context"
	"sencha-twirp-rpc/internal/models"
)

type Backender interface {
	GetTheme(ctx context.Context, name string) (*models.Theme, error)
}
