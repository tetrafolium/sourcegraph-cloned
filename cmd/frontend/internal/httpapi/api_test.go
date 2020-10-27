package httpapi

import (
	"github.com/gorilla/mux"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/enterprise"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/httpapi/router"
	"github.com/tetrafolium/sourcegraph-cloned/internal/httptestutil"
	"github.com/tetrafolium/sourcegraph-cloned/internal/txemail"
)

func init() {
	txemail.DisableSilently()
}

func newTest() *httptestutil.Client {
	enterpriseServices := enterprise.DefaultServices()

	return httptestutil.NewTest(NewHandler(
		router.New(mux.NewRouter()),
		nil,
		enterpriseServices.GitHubWebhook,
		enterpriseServices.GitLabWebhook,
		enterpriseServices.BitbucketServerWebhook,
		enterpriseServices.NewCodeIntelUploadHandler,
	))
}
