package graphqlbackend

import (
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/graphqlbackend"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/internal/dotcom/billing"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/internal/dotcom/productsubscription"
)

func init() {
	// TODO(efritz) - de-globalize assignments in this function
	// Contribute the GraphQL types DotcomMutation and DotcomQuery.
	graphqlbackend.Dotcom = dotcomResolver{}
}

// dotcomResolver implements the GraphQL types DotcomMutation and DotcomQuery.
type dotcomResolver struct {
	productsubscription.ProductSubscriptionLicensingResolver
	billing.BillingResolver
}
