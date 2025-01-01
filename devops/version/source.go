package version

import (
	"context"
	"fmt"
	"strings"

	"github.com/unmango/go/fs/github"
	"github.com/unmango/go/lazy"
)

type Source interface {
	Latest(context.Context) (string, error)
	Name(context.Context) (string, error)
}

func GuessSource(s string) (Source, error) {
	switch {
	case strings.Contains("github", s):
		return GitHub(s), nil
	default:
		return nil, fmt.Errorf("unrecognized source: %s", s)
	}
}

type GitHub string

var ghclient = lazy.Once(func() *github.Client {
	return github.NewClient(nil)
})

// Name implements Source.
func (g GitHub) Name(context.Context) (string, error) {
	panic("unimplemented")
}

// Latest implements Source.
func (g GitHub) Latest(context.Context) (string, error) {
	panic("unimplemented")
}
