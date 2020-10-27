package app

import (
	"net/http"

	"github.com/inconshreveable/log15"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/backend"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/usagestats"
)

func usageStatsArchiveHandler(w http.ResponseWriter, r *http.Request) {
	// 🚨SECURITY: Only site admins may get this archive.
	if err := backend.CheckCurrentUserIsSiteAdmin(r.Context()); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=\"SourcegraphUsersUsageArchive.zip\"")

	archive, err := usagestats.GetArchive(r.Context())
	if err != nil {
		log15.Error("usagestats.WriteArchive", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(archive)
}
