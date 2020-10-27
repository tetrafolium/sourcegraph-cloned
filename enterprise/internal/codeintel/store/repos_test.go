package store

import (
	"context"
	"testing"

	"github.com/tetrafolium/sourcegraph-cloned/internal/db/dbconn"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db/dbtesting"
)

func TestRepoName(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	dbtesting.SetupGlobalTestDB(t)
	store := testStore()

	if _, err := dbconn.Global.Exec(`INSERT INTO repo (id, name) VALUES (50, 'github.com/foo/bar')`); err != nil {
		t.Fatalf("unexpected error inserting repo: %s", err)
	}

	name, err := store.RepoName(context.Background(), 50)
	if err != nil {
		t.Fatalf("unexpected error getting repo name: %s", err)
	}
	if name != "github.com/foo/bar" {
		t.Errorf("unexpected repo name. want=%s have=%s", "github.com/foo/bar", name)
	}
}
