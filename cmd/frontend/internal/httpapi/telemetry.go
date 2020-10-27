package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/inconshreveable/log15"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/usagestats"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/usagestatsdeprecated"
	"github.com/tetrafolium/sourcegraph-cloned/internal/eventlogger"
)

var telemetryHandler http.Handler

func init() {
	telemetryHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tr eventlogger.TelemetryRequest
		err := json.NewDecoder(r.Body).Decode(&tr)
		if err != nil {
			log15.Error("telemetryHandler: Decode", "error", err)
		}
		err = usagestats.LogBackendEvent(tr.UserID, tr.EventName, tr.Argument)
		if err != nil {
			log15.Error("telemetryHandler: usagestats.LogBackendEvent", "error", err)
		}
		if tr.UserID != 0 && tr.EventName == "SavedSearchEmailNotificationSent" {
			err = usagestatsdeprecated.LogActivity(true, tr.UserID, "", "STAGEVERIFY")
			if err != nil {
				log15.Error("telemetryHandler: usagestats.LogBackendEvent", "error", err)
			}
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
