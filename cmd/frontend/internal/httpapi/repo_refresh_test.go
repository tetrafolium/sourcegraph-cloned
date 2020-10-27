package httpapi

import (
	"context"
	"testing"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/backend"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/types"
	"github.com/tetrafolium/sourcegraph-cloned/internal/api"
	"github.com/tetrafolium/sourcegraph-cloned/internal/gitserver"
	"github.com/tetrafolium/sourcegraph-cloned/internal/repoupdater"
	"github.com/tetrafolium/sourcegraph-cloned/internal/repoupdater/protocol"
)

func TestRepoRefresh(t *testing.T) {
	c := newTest()

	enqueueRepoUpdateCount := map[api.RepoName]int{}
	repoupdater.MockEnqueueRepoUpdate = func(ctx context.Context, repo gitserver.Repo) (*protocol.RepoUpdateResponse, error) {
		if exp := "git@github.com:dummy-url"; repo.URL != exp {
			t.Errorf("missing or incorrect clone URL, expected %q, got %q", exp, repo.URL)
		}
		enqueueRepoUpdateCount[repo.Name]++
		return nil, nil
	}
	repoupdater.MockRepoLookup = func(args protocol.RepoLookupArgs) (*protocol.RepoLookupResult, error) {
		return &protocol.RepoLookupResult{
			Repo: &protocol.RepoInfo{
				Name: args.Repo,
				VCS: protocol.VCSInfo{
					URL: "git@github.com:dummy-url",
				},
			},
		}, nil
	}

	backend.Mocks.Repos.GetByName = func(ctx context.Context, name api.RepoName) (*types.Repo, error) {
		switch name {
		case "github.com/gorilla/mux":
			return &types.Repo{ID: 2, Name: name}, nil
		default:
			panic("wrong path")
		}
	}
	backend.Mocks.Repos.ResolveRev = func(ctx context.Context, repo *types.Repo, rev string) (api.CommitID, error) {
		if repo.ID != 2 || rev != "master" {
			t.Error("wrong arguments to ResolveRev")
		}
		return "aed", nil
	}

	if _, err := c.PostOK("/repos/github.com/gorilla/mux/-/refresh", nil); err != nil {
		t.Fatal(err)
	}
	if ct := enqueueRepoUpdateCount["github.com/gorilla/mux"]; ct != 1 {
		t.Errorf("expected EnqueueRepoUpdate to be called once, but was called %d times", ct)
	}
}
