package httpapi

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/handlerutil"
	"github.com/tetrafolium/sourcegraph-cloned/internal/gitserver"
	"github.com/tetrafolium/sourcegraph-cloned/internal/repoupdater"
	"github.com/tetrafolium/sourcegraph-cloned/internal/repoupdater/protocol"
)

func serveRepoRefresh(w http.ResponseWriter, r *http.Request) error {
	repo, err := handlerutil.GetRepo(r.Context(), mux.Vars(r))
	if err != nil {
		return err
	}
	repoMeta, err := repoupdater.DefaultClient.RepoLookup(context.Background(), protocol.RepoLookupArgs{
		Repo: repo.Name,
	})
	if err != nil {
		return err
	}
	_, err = repoupdater.DefaultClient.EnqueueRepoUpdate(context.Background(), gitserver.Repo{
		Name: repo.Name,
		URL:  repoMeta.Repo.VCS.URL,
	})
	return err
}
