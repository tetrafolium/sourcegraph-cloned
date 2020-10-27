package gitlaboauth

import (
	"net/http"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/auth"
	"github.com/tetrafolium/sourcegraph-cloned/enterprise/cmd/frontend/auth/oauth"
	"github.com/tetrafolium/sourcegraph-cloned/internal/extsvc"
	"github.com/tetrafolium/sourcegraph-cloned/schema"
)

const authPrefix = auth.AuthURLPrefix + "/gitlab"

func init() {
	oauth.AddIsOAuth(func(p schema.AuthProviders) bool {
		return p.Gitlab != nil
	})
}

var Middleware = &auth.Middleware{
	API: func(next http.Handler) http.Handler {
		return oauth.NewHandler(extsvc.TypeGitLab, authPrefix, true, next)
	},
	App: func(next http.Handler) http.Handler {
		return oauth.NewHandler(extsvc.TypeGitLab, authPrefix, false, next)
	},
}
