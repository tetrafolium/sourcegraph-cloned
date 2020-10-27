package cli

import (
	"reflect"
	"testing"

	"github.com/tetrafolium/sourcegraph-cloned/internal/conf/conftypes"
)

func TestServiceConnections(t *testing.T) {
	// We only test that we get something non-empty back.
	sc := serviceConnections()
	if reflect.DeepEqual(sc, conftypes.ServiceConnections{}) {
		t.Fatal("expected non-empty service connections")
	}
}
