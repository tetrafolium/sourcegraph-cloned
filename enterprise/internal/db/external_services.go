package db

import (
	"github.com/tetrafolium/sourcegraph-cloned/internal/authz/bitbucketserver"
	"github.com/tetrafolium/sourcegraph-cloned/internal/authz/github"
	"github.com/tetrafolium/sourcegraph-cloned/internal/authz/gitlab"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db"
	"github.com/tetrafolium/sourcegraph-cloned/schema"
)

// NewExternalServicesStore returns an OSS db.ExternalServicesStore set with
// enterprise validators.
func NewExternalServicesStore() *db.ExternalServicesStore {
	return &db.ExternalServicesStore{
		GitHubValidators: []func(*schema.GitHubConnection) error{
			github.ValidateAuthz,
		},
		GitLabValidators: []func(*schema.GitLabConnection, []schema.AuthProviders) error{
			gitlab.ValidateAuthz,
		},
		BitbucketServerValidators: []func(*schema.BitbucketServerConnection) error{
			bitbucketserver.ValidateAuthz,
		},
	}
}
