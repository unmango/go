package release

import "github.com/google/go-github/v66/github"

type state struct {
	release *github.RepositoryRelease
	assets  map[int64]*github.ReleaseAsset
}
