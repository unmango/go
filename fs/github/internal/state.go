package internal

import (
	"github.com/google/go-github/v66/github"
)

type State struct {
	Users    map[string]*github.User
	Repos    map[int64]*github.Repository
	Releases map[int64]*github.RepositoryRelease
	Assets   map[int64]*github.ReleaseAsset
}
