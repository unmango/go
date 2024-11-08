package github

import (
	"net/http"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
	repo "github.com/unmango/go/fs/github/repository"
	"github.com/unmango/go/option"
)

type GitHub struct {
	afero.Fs
	owner, repo string

	Client *github.Client
}

var _ afero.Fs = &GitHub{}

func DefaultClient() *github.Client {
	return github.NewClient(http.DefaultClient)
}

func New(owner, repository string, options ...Option) *GitHub {
	gh := &GitHub{
		owner: owner,
		repo:  repository,
	}
	option.ApplyAll(gh, options)

	if gh.Client == nil {
		gh.Client = DefaultClient()
	}
	if gh.Fs == nil {
		gh.Fs = repo.New(
			owner, repository,
			gh.Client.Repositories,
		)
	}

	return gh
}
