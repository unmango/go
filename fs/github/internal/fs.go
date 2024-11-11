package internal

import "github.com/google/go-github/v66/github"

type Fs struct {
	ContextAccessor
	Client *github.Client
}
