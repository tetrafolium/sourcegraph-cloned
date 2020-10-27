// Package shared contains the frontend command implementation shared
package shared

import (
	"fmt"
	"os"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/enterprise"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/cli"
	"github.com/tetrafolium/sourcegraph-cloned/internal/env"

	_ "github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/registry"
)

// Main is the main function that runs the frontend process.
//
// It is exposed as function in a package so that it can be called by other
// main package implementations such as Sourcegraph Enterprise, which import
// proprietary/private code.
func Main(enterpriseSetupHook func() enterprise.Services) {
	env.Lock()
	err := cli.Main(enterpriseSetupHook)
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal:", err)
		os.Exit(1)
	}
}
