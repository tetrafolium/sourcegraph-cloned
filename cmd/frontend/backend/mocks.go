package backend

import (
	"context"

	"github.com/tetrafolium/sourcegraph-cloned/internal/actor"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db"
	"github.com/tetrafolium/sourcegraph-cloned/internal/trace/ot"
	"github.com/tetrafolium/sourcegraph-cloned/internal/vcs/git"
)

var Mocks MockServices

type MockServices struct {
	Repos MockRepos
}

// testContext creates a new context.Context for use by tests
func testContext() context.Context {
	db.Mocks = db.MockStores{}
	Mocks = MockServices{}
	git.ResetMocks()

	ctx := context.Background()
	ctx = actor.WithActor(ctx, &actor.Actor{UID: 1})
	_, ctx = ot.StartSpanFromContext(ctx, "dummy")

	return ctx
}
