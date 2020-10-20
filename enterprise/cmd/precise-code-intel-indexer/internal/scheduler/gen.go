package scheduler

//go:generate env GOBIN=$PWD/.bin GO111MODULE=on go install github.com/efritz/go-mockgen
//go:generate $PWD/.bin/go-mockgen -f github.com/tetrafolium/sourcegraph/enterprise/cmd/precise-code-intel-indexer/internal/scheduler -i gitserverClient -o mock_gitserver_client_test.go
