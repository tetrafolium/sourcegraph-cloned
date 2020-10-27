package httpheader

import (
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/auth/providers"
	"github.com/tetrafolium/sourcegraph-cloned/internal/conf"
)

// Watch for configuration changes related to the http-header auth provider.
func init() {
	go func() {
		conf.Watch(func() {
			newPC, _ := getProviderConfig()
			if newPC == nil {
				providers.Update("httpheader", nil)
				return
			}
			providers.Update("httpheader", []providers.Provider{&provider{c: newPC}})
		})
	}()
}
