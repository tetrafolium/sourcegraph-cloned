package mocks

//go:generate env GOBIN=$PWD/.bin GO111MODULE=on go install github.com/efritz/go-mockgen
//go:generate $PWD/.bin/go-mockgen -f github.com/tetrafolium/sourcegraph-cloned/enterprise/internal/codeintel/bundles/persistence -i Store -o mock_store.go
