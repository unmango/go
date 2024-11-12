package internal

import (
	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/lazy"
	"github.com/unmango/go/result"
)

type Fs struct {
	ContextAccessor
	Client *github.Client
}

type Result[T any] result.Result2[T, *github.Response]

func Request[T any](req Result[T]) lazy.Lazy[Result[T]] {
	return lazy.Once(lazy.Of(req))
}
