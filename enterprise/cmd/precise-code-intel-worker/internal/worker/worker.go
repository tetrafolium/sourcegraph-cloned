package worker

import (
	"context"
	"database/sql"
	"time"

	"github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/precise-code-intel-worker/internal/metrics"
	bundles "github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/bundles/client"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/bundles/persistence"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/bundles/persistence/postgres"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/store"
	"github.com/tetrafolium/sourcegraph-cloned/internal/actor"
	"github.com/tetrafolium/sourcegraph-cloned/internal/observation"
	"github.com/tetrafolium/sourcegraph-cloned/internal/workerutil"
	"github.com/tetrafolium/sourcegraph-cloned/internal/workerutil/dbworker"
)

func NewWorker(
	s store.Store,
	codeIntelDB *sql.DB,
	bundleManagerClient bundles.BundleManagerClient,
	gitserverClient gitserverClient,
	pollInterval time.Duration,
	numProcessorRoutines int,
	budgetMax int64,
	metrics metrics.WorkerMetrics,
	observationContext *observation.Context,
) *workerutil.Worker {
	rootContext := actor.WithActor(context.Background(), &actor.Actor{Internal: true})

	handler := &handler{
		store:               s,
		bundleManagerClient: bundleManagerClient,
		gitserverClient:     gitserverClient,
		metrics:             metrics,
		enableBudget:        budgetMax > 0,
		budgetRemaining:     budgetMax,
		createStore: func(id int) persistence.Store {
			return persistence.NewObserved(postgres.NewStore(codeIntelDB, id), observationContext)
		},
	}

	return dbworker.NewWorker(rootContext, store.WorkerutilUploadStore(s), handler, workerutil.WorkerOptions{
		NumHandlers: numProcessorRoutines,
		Interval:    pollInterval,
		Metrics: workerutil.WorkerMetrics{
			HandleOperation: metrics.ProcessOperation,
		},
	})
}
