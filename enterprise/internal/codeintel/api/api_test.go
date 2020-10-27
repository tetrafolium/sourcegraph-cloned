package api

import (
	bundles "github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/bundles/client"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/store"
	"github.com/tetrafolium/sourcegraph-cloned/internal/observation"
)

func testAPI(store store.Store, bundleManagerClient bundles.BundleManagerClient, gitserverClient gitserverClient) CodeIntelAPI {
	// Wrap in observed, as that's how it's used in production
	return NewObserved(New(store, bundleManagerClient, gitserverClient), &observation.TestContext)
}
