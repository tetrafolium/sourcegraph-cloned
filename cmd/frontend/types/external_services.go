package types

import (
	"github.com/tetrafolium/sourcegraph-cloned/schema"
)

type BitbucketServerConnection struct {
	// The unique resource identifier of the external service.
	URN string
	*schema.BitbucketServerConnection
}

type GitHubConnection struct {
	// The unique resource identifier of the external service.
	URN string
	*schema.GitHubConnection
}

type GitLabConnection struct {
	// The unique resource identifier of the external service.
	URN string
	*schema.GitLabConnection
}
