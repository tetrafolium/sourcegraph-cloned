// Command frontend contains the enterprise frontend implementation.
//
// It wraps the open source frontend command and merely injects a few
// proprietary things into it via e.g. blank/underscore imports in this file
// which register side effects with the frontend package.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/enterprise"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/shared"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/internal/authz"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/internal/campaigns"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/internal/codeintel"
	licensing "github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/internal/licensing/init"

	_ "github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/auth"
	_ "github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/internal/graphqlbackend"
	_ "github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/internal/registry"
)

func main() {
	shared.Main(enterpriseSetupHook)
}

var initFunctions = map[string]func(ctx context.Context, enterpriseServices *enterprise.Services) error{
	"authz":     authz.Init,
	"campaigns": campaigns.Init,
	"codeintel": codeintel.Init,
	"licensing": licensing.Init,
}

func enterpriseSetupHook() enterprise.Services {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if debug {
		log.Println("enterprise edition")
	}

	ctx := context.Background()
	enterpriseServices := enterprise.DefaultServices()

	for name, fn := range initFunctions {
		if err := fn(ctx, &enterpriseServices); err != nil {
			log.Fatal(fmt.Sprintf("failed to initialize %s: %s", name, err))
		}
	}

	return enterpriseServices
}
