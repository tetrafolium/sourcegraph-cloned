package backend

import (
	"context"

	"github.com/tetrafolium/sourcegraph-cloned/internal/search"
	symbolsclient "github.com/tetrafolium/sourcegraph-cloned/internal/symbols"
	"github.com/tetrafolium/sourcegraph-cloned/internal/symbols/protocol"
)

// Symbols backend.
var Symbols = &symbols{}

type symbols struct{}

// ListTags returns symbols in a repository from ctags.
func (symbols) ListTags(ctx context.Context, args search.SymbolsParameters) ([]protocol.Symbol, error) {
	result, err := symbolsclient.DefaultClient.Search(ctx, args)
	if result == nil {
		return nil, err
	}
	return result.Symbols, err
}
