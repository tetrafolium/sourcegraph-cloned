package saml

import (
	"context"

	"github.com/inconshreveable/log15"
	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/auth/providers"
	"github.com/tetrafolium/sourcegraph-cloned/internal/conf"
	"github.com/tetrafolium/sourcegraph-cloned/schema"
)

func getProviders() []providers.Provider {
	var cfgs []*schema.SAMLAuthProvider
	for _, p := range conf.Get().AuthProviders {
		if p.Saml == nil {
			continue
		}
		cfgs = append(cfgs, withConfigDefaults(p.Saml))
	}
	multiple := len(cfgs) >= 2
	ps := make([]providers.Provider, 0, len(cfgs))
	for _, cfg := range cfgs {
		p := &provider{config: *cfg, multiple: multiple}
		ps = append(ps, p)
	}
	return ps
}

func init() {
	go func() {
		conf.Watch(func() {
			ps := getProviders()
			for _, p := range ps {
				go func(p providers.Provider) {
					if err := p.Refresh(context.Background()); err != nil {
						log15.Error("Error prefetching SAML service provider metadata.", "error", err)
					}
				}(p)
			}
			providers.Update("saml", ps)
		})
	}()
}
