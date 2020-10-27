package graphql

import (
	"context"

	gql "github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/graphqlbackend"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/graphqlbackend/graphqlutil"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/resolvers"
)

type DiagnosticConnectionResolver struct {
	diagnostics      []resolvers.AdjustedDiagnostic
	totalCount       int
	locationResolver *CachedLocationResolver
}

func NewDiagnosticConnectionResolver(diagnostics []resolvers.AdjustedDiagnostic, totalCount int, locationResolver *CachedLocationResolver) gql.DiagnosticConnectionResolver {
	return &DiagnosticConnectionResolver{
		diagnostics:      diagnostics,
		totalCount:       totalCount,
		locationResolver: locationResolver,
	}
}

func (r *DiagnosticConnectionResolver) Nodes(ctx context.Context) ([]gql.DiagnosticResolver, error) {
	resolvers := make([]gql.DiagnosticResolver, 0, len(r.diagnostics))
	for i := range r.diagnostics {
		resolvers = append(resolvers, NewDiagnosticResolver(r.diagnostics[i], r.locationResolver))
	}
	return resolvers, nil
}

func (r *DiagnosticConnectionResolver) TotalCount(ctx context.Context) (int32, error) {
	return int32(r.totalCount), nil
}

func (r *DiagnosticConnectionResolver) PageInfo(ctx context.Context) (*graphqlutil.PageInfo, error) {
	return graphqlutil.HasNextPage(len(r.diagnostics) < r.totalCount), nil
}
