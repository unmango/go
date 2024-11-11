package github

import (
	"io/fs"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github/internal"
	"github.com/unmango/go/option"
)

type Fs struct {
	afero.ReadOnlyFs
	internal.ContextAccessor

	Client *github.Client
	users  map[string]*github.User
}

var _ afero.Fs = &Fs{}

// Name implements afero.Fs.
func (g *Fs) Name() string {
	return "https://github.com"
}

// Open implements afero.Fs.
func (g *Fs) Open(name string) (afero.File, error) {
	if _, ok := g.users[name]; !ok {
		if u, err := g.getUser(name); err != nil {
			return nil, err
		} else {
			g.users[name] = u
		}
	}

	return &userFile{g, g.users[name]}, nil
}

// OpenFile implements afero.Fs.
func (g *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return g.Open(name)
}

// Stat implements afero.Fs.
func (g *Fs) Stat(name string) (fs.FileInfo, error) {
	panic("unimplemented")
}

func (g *Fs) getUser(name string) (*github.User, error) {
	user, _, err := g.Client.Users.Get(g.Context(), name)
	return user, err
}

func New(owner, repository string, options ...Option) *Fs {
	gh := &Fs{}
	option.ApplyAll(gh, options)

	if gh.Client == nil {
		gh.Client = DefaultClient()
	}

	return gh
}
