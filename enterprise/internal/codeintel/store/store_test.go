package store

import (
	"database/sql"

	"github.com/tetrafolium/sourcegraph-cloned/internal/db/basestore"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db/dbconn"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db/dbtesting"
	"github.com/tetrafolium/sourcegraph-cloned/internal/observation"
)

func init() {
	dbtesting.DBNameSuffix = "codeintel"
}

func testStore() Store {
	// Wrap in observed, as that's how it's used in production
	return NewObserved(&store{Store: basestore.NewWithDB(dbconn.Global, sql.TxOptions{})}, &observation.TestContext)
}
