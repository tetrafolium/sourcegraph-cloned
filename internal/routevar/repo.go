package routevar

import (
	"strings"

	"github.com/tetrafolium/sourcegraph-cloned/internal/api"
	"github.com/tetrafolium/sourcegraph-cloned/internal/lazyregexp"
)

// A RepoRev specifies a repo at a revision. The revision need not be an absolute
// commit ID. This RepoRev type is appropriate for user input (e.g.,
// from a URL), where it is convenient to allow users to specify
// non-absolute commit IDs that the server can resolve.
type RepoRev struct {
	Repo api.RepoName // a repo path
	Rev  string       // a VCS revision specifier (branch, "master~7", commit ID, etc.)
}

var (
	Repo = `{Repo:` + namedToNonCapturingGroups(RepoPattern) + `}`
	Rev  = `{Rev:` + namedToNonCapturingGroups(RevPattern) + `}`

	RepoRevSuffix = `{Rev:` + namedToNonCapturingGroups(`(?:@`+RevPattern+`)?`) + `}`
)

const (
	// RepoPattern is the regexp pattern that matches repo path strings
	// ("repo" or "domain.com/repo" or "domain.com/path/to/repo").
	RepoPattern = `(?P<repo>(?:` + pathComponentNotDelim + `/)*` + pathComponentNotDelim + `)`

	RepoPathDelim         = "-"
	pathComponentNotDelim = `(?:[^@/` + RepoPathDelim + `]|(?:[^/@]{2,}))`

	// RevPattern is the regexp pattern that matches a VCS revision
	// specifier (e.g., "master" or "my/branch~1", or a full 40-char
	// commit ID).
	RevPattern = `(?P<rev>(?:` + pathComponentNotDelim + `/)*` + pathComponentNotDelim + `)`
)

var repoPattern = lazyregexp.New("^" + RepoPattern + "$")

// ParseRepo parses a repo path string. If spec is invalid, an
// InvalidError is returned.
func ParseRepo(spec string) (repo api.RepoName, err error) {
	if m := repoPattern.FindStringSubmatch(spec); len(m) > 0 {
		repo = api.RepoName(m[0])
		return
	}
	return "", InvalidError{"Repo", spec, nil}
}

// RepoRouteVars returns route variables for constructing repository
// routes.
func RepoRouteVars(repo api.RepoName) map[string]string {
	return map[string]string{"Repo": string(repo)}
}

// ToRepoRev marshals a map containing route variables
// generated by (RepoRevSpec).RouteVars() and returns the equivalent
// RepoRevSpec struct.
func ToRepoRev(routeVars map[string]string) RepoRev {
	rr := RepoRev{Repo: ToRepo(routeVars)}
	if revStr := routeVars["Rev"]; revStr != "" {
		if !strings.HasPrefix(revStr, "@") {
			panic("Rev should have had '@' prefix from route")
		}
		rr.Rev = strings.TrimPrefix(revStr, "@")
	}
	if _, ok := routeVars["CommitID"]; ok {
		panic("unexpected CommitID route var; was removed in the simple-routes branch")
	}
	return rr
}

// ToRepo returns the repo path string from a map containing route variables.
func ToRepo(routeVars map[string]string) api.RepoName {
	return api.RepoName(routeVars["Repo"])
}

// RepoRevRouteVars returns route variables for constructing routes to a
// repository revision.
func RepoRevRouteVars(s RepoRev) map[string]string {
	m := RepoRouteVars(s.Repo)
	var rev string
	if s.Rev != "" {
		rev = "@" + s.Rev
	}
	m["Rev"] = rev
	return m
}
