package server

import (
	indexmanager "github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/precise-code-intel-indexer/internal/index_manager"
	"github.com/tetrafolium/sourcegraph-cloned/internal/goroutine"
	"github.com/tetrafolium/sourcegraph-cloned/internal/httpserver"
)

const Port = 3189

type Server struct {
	indexManager indexmanager.Manager
}

func New(indexManager indexmanager.Manager) goroutine.BackgroundRoutine {
	server := &Server{
		indexManager: indexManager,
	}

	return httpserver.New(Port, server.setupRoutes)
}
