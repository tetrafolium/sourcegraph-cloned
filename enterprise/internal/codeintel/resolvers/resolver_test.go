package resolvers

import (
	"context"
	"testing"

	gql "github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/graphqlbackend"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/types"
	apimocks "github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/api/mocks"
	bundlemocks "github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/bundles/client/mocks"
	storemocks "github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/store/mocks"
	"github.com/tetrafolium/sourcegraph-cloned/internal/api"
)

func TestQueryResolver(t *testing.T) {
	mockStore := storemocks.NewMockStore()
	mockBundleManagerClient := bundlemocks.NewMockBundleManagerClient()
	mockCodeIntelAPI := apimocks.NewMockCodeIntelAPI() // returns no dumps

	resolver := NewResolver(mockStore, mockBundleManagerClient, mockCodeIntelAPI, nil)
	queryResolver, err := resolver.QueryResolver(context.Background(), &gql.GitBlobLSIFDataArgs{
		Repo:      &types.Repo{ID: 50},
		Commit:    api.CommitID("deadbeef"),
		Path:      "/foo/bar.go",
		ExactPath: true,
		ToolName:  "lsif-go",
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if queryResolver != nil {
		t.Errorf("expected nil-valued resolver")
	}
}
