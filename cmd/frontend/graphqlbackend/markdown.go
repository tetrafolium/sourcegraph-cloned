package graphqlbackend

import "github.com/tetrafolium/sourcegraph-cloned/internal/markdown"

type Markdown string

func (m Markdown) Text() string {
	return string(m)
}

func (m Markdown) HTML() string {
	return markdown.Render(string(m))
}
