package github

import (
	"net/url"

	"github.com/tetrafolium/sourcegraph-cloned/internal/api"
	"github.com/tetrafolium/sourcegraph-cloned/internal/extsvc"
)

// ExternalRepoSpec returns an api.ExternalRepoSpec that refers to the specified GitHub repository.
func ExternalRepoSpec(repo *Repository, baseURL url.URL) api.ExternalRepoSpec {
	return api.ExternalRepoSpec{
		ID:          repo.ID,
		ServiceType: extsvc.TypeGitHub,
		ServiceID:   extsvc.NormalizeBaseURL(&baseURL).String(),
	}
}
