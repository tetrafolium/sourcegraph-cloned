package db

import (
	"context"

	"github.com/tetrafolium/sourcegraph-cloned/internal/api"
)

type MockSettings struct {
	GetLatest        func(ctx context.Context, subject api.SettingsSubject) (*api.Settings, error)
	CreateIfUpToDate func(ctx context.Context, subject api.SettingsSubject, lastID, authorUserID *int32, contents string) (latestSetting *api.Settings, err error)
}
