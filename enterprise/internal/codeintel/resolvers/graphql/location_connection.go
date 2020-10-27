package graphql

import (
	"context"

	gql "github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/graphqlbackend"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/graphqlbackend/graphqlutil"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/resolvers"
)

type LocationConnectionResolver struct {
	locations        []resolvers.AdjustedLocation
	cursor           *string
	locationResolver *CachedLocationResolver
}

func NewLocationConnectionResolver(locations []resolvers.AdjustedLocation, cursor *string, locationResolver *CachedLocationResolver) gql.LocationConnectionResolver {
	return &LocationConnectionResolver{
		locations:        locations,
		cursor:           cursor,
		locationResolver: locationResolver,
	}
}

func (r *LocationConnectionResolver) Nodes(ctx context.Context) ([]gql.LocationResolver, error) {
	return resolveLocations(ctx, r.locationResolver, r.locations)
}

func (r *LocationConnectionResolver) PageInfo(ctx context.Context) (*graphqlutil.PageInfo, error) {
	return encodeCursor(r.cursor), nil
}
