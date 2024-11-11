package github

import "github.com/google/go-github/v66/github"

type repositoryFile struct {
	*Fs
	repo *github.Repository
}
