package sqlite

import "github.com/tetrafolium/sourcegraph-cloned/internal/sqliteutil"

func init() {
	sqliteutil.SetLocalLibpath()
	sqliteutil.MustRegisterSqlite3WithPcre()
}
