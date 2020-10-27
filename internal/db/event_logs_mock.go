package db

import (
	"context"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/types"
)

type MockEventLogs struct {
	LatestPing func(ctx context.Context) (*types.Event, error)
}
