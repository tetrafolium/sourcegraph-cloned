package campaigns

import (
	"context"
	"database/sql"
	"time"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/enterprise"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/repo-updater/repos"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/campaigns"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/campaigns/resolvers"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db/dbconn"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db/globalstatedb"
)

func Init(ctx context.Context, enterpriseServices *enterprise.Services) error {
	globalState, err := globalstatedb.Get(ctx)
	if err != nil {
		return err
	}

	campaignsStore := campaigns.NewStoreWithClock(dbconn.Global, msResolutionClock)
	repositories := repos.NewDBStore(dbconn.Global, sql.TxOptions{})

	enterpriseServices.CampaignsResolver = resolvers.NewResolver(dbconn.Global)
	enterpriseServices.GitHubWebhook = campaigns.NewGitHubWebhook(campaignsStore, repositories, msResolutionClock)
	enterpriseServices.BitbucketServerWebhook = campaigns.NewBitbucketServerWebhook(
		campaignsStore,
		repositories,
		msResolutionClock,
		"sourcegraph-"+globalState.SiteID,
	)
	enterpriseServices.GitLabWebhook = campaigns.NewGitLabWebhook(campaignsStore, repositories, msResolutionClock)

	return nil
}

var msResolutionClock = func() time.Time { return time.Now().UTC().Truncate(time.Microsecond) }
