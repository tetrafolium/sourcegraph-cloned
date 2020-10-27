package externallink

import (
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/backend"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db"
	"github.com/tetrafolium/sourcegraph-cloned/internal/repoupdater"
)

func resetMocks() {
	repoupdater.MockRepoLookup = nil
	db.Mocks = db.MockStores{}
	backend.Mocks = backend.MockServices{}
}
