package router

import (
	"net/http"
	"path"
	"strings"

	"github.com/tetrafolium/sourcegraph-cloned/internal/api"
	"github.com/tetrafolium/sourcegraph-cloned/internal/routevar"

	"github.com/gorilla/mux"
)

// same as spec.unresolvedRevPattern but also not allowing path
// components starting with ".".
const revSuffixNoDots = `{Rev:(?:@(?:(?:[^@=/.-]|(?:[^=/@.]{2,}))/)*(?:[^@=/.-]|(?:[^=/@.]{2,})))?}`

func addOldTreeRedirectRoute(matchRouter *mux.Router) {
	matchRouter.Path("/" + routevar.Repo + revSuffixNoDots + `/.tree{Path:.*}`).Methods("GET").Name(OldTreeRedirect).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		path := path.Clean(v["Path"])
		if !strings.HasPrefix(path, "/") && path != "" {
			path = "/" + path
		}

		http.Redirect(w, r, URLToRepoTreeEntry(api.RepoName(v["Repo"]), v["Rev"], path).String(), http.StatusMovedPermanently)
	})
}
