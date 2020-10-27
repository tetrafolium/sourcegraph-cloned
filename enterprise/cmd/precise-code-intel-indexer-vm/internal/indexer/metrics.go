package indexer

import (
	"github.com/tetrafolium/sourcegraph-cloned/internal/metrics"
	"github.com/tetrafolium/sourcegraph-cloned/internal/observation"
)

type IndexerMetrics struct {
	ProcessOperation *observation.Operation
}

func NewIndexerMetrics(observationContext *observation.Context) IndexerMetrics {
	metrics := metrics.NewOperationMetrics(
		observationContext.Registerer,
		"index_queue_processor",
		metrics.WithLabels("op"),
		metrics.WithCountHelp("Total number of records processed"),
	)

	return IndexerMetrics{
		ProcessOperation: observationContext.Operation(observation.Op{
			Name:         "Processor.Process",
			MetricLabels: []string{"process"},
			Metrics:      metrics,
		}),
	}
}
