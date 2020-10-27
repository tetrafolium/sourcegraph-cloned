package auth

import (
	"testing"

	"github.com/tetrafolium/sourcegraph-cloned/internal/conf"
	"github.com/tetrafolium/sourcegraph-cloned/schema"
)

func TestValidateCustom(t *testing.T) {
	tests := map[string]struct {
		input        conf.Unified
		wantProblems conf.Problems
	}{
		"no auth.providers": {
			input:        conf.Unified{SiteConfiguration: schema.SiteConfiguration{}},
			wantProblems: conf.NewSiteProblems("no auth providers set"),
		},
		"empty auth.providers": {
			input:        conf.Unified{SiteConfiguration: schema.SiteConfiguration{AuthProviders: []schema.AuthProviders{}}},
			wantProblems: conf.NewSiteProblems("no auth providers set"),
		},
		"single auth provider": {
			input: conf.Unified{SiteConfiguration: schema.SiteConfiguration{
				AuthProviders: []schema.AuthProviders{
					{Builtin: &schema.BuiltinAuthProvider{Type: "a"}},
				},
			}},
			wantProblems: nil,
		},
		"multiple auth providers": {
			input: conf.Unified{SiteConfiguration: schema.SiteConfiguration{
				AuthProviders: []schema.AuthProviders{
					{Builtin: &schema.BuiltinAuthProvider{Type: "a"}},
					{Builtin: &schema.BuiltinAuthProvider{Type: "b"}},
				},
			}},
			wantProblems: nil,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			conf.TestValidator(t, test.input, validateConfig, test.wantProblems)
		})
	}
}
