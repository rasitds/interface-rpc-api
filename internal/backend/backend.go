package backend

import "context"

type Backender interface {
	GetColors(ctx context.Context, name string) (string, error)
}
