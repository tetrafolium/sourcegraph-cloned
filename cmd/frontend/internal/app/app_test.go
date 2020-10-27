package app

import "github.com/tetrafolium/sourcegraph-cloned/internal/txemail"

func init() {
	txemail.DisableSilently()
}
