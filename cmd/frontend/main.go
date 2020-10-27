package main

import (
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/enterprise"
	_ "github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/app/assets"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/shared"
	"github.com/tetrafolium/sourcegraph-cloned/internal/authz"
)

// Note: All frontend code should be added to shared.Main, not here. See that
// function for details.

func main() {
	// Set dummy authz provider to unblock channel for checking permissions in GraphQL APIs.
	// See https://github.com/tetrafolium/sourcegraph-cloned/issues/3847 for details.
	authz.SetProviders(true, []authz.Provider{})

	shared.Main(func() enterprise.Services {
		return enterprise.DefaultServices()
	})
}
