package github

import "github.com/google/go-github/v66/github"

type Option func(*GitHub)

func WithClient(client *github.Client) Option {
	return func(gh *GitHub) {
		gh.Client = client
	}
}

func WithRepository(owner, repository string) Option {
	return func(gh *GitHub) {
		gh.owner = owner
		gh.repo = repository
	}
}
