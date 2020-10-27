package resetter

import (
	"time"

	"github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/store"
	"github.com/tetrafolium/sourcegraph-cloned/internal/workerutil/dbworker"
)

func NewUploadResetter(
	s store.Store,
	resetInterval time.Duration,
	metrics dbworker.ResetterMetrics,
) *dbworker.Resetter {
	return dbworker.NewResetter(store.WorkerutilUploadStore(s), dbworker.ResetterOptions{
		Name:     "upload resetter",
		Interval: resetInterval,
		Metrics:  metrics,
	})
}
