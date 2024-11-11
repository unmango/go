package github

import (
	"context"

	"github.com/google/go-github/v66/github"
)

type Option func(*Fs)

func WithClient(client *github.Client) Option {
	return func(gh *Fs) {
		gh.Client = client
	}
}

func WithContext(ctx context.Context) Option {
	return func(f *Fs) {
		f.ContextAccessor = func() context.Context {
			return ctx
		}
	}
}
