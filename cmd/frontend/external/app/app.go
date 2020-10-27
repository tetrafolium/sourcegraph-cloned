// Package app exports symbols from frontend/internal/app. See the parent
// package godoc for more information.
package app

import (
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/app"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/app/jscontext"
)

type SignOutURL = app.SignOutURL

var RegisterSSOSignOutHandler = app.RegisterSSOSignOutHandler

func SetBillingPublishableKey(value string) {
	jscontext.BillingPublishableKey = value
}
