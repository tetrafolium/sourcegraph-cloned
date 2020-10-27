package bg

import (
	"github.com/tetrafolium/sourcegraph-cloned/internal/db/dbtesting"
)

func init() {
	dbtesting.DBNameSuffix = "bgdb"
}
