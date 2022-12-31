package postgresql

import (
	"context"
	"sencha-twirp-rpc/internal/backend"
)

type PostgreSQLBackend struct{}

func NewPostgreSQLBackend() backend.Backender {
	return &PostgreSQLBackend{}
}

func (b *PostgreSQLBackend) GetColors(ctx context.Context, name string) (string, error) {
	return name, nil
}
