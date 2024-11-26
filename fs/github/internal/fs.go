package internal

import (
	"context"
	"fmt"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
)

func Open(ctx context.Context, gh *github.Client, path Path) (afero.File, error) {
	owner, err := path.Owner()
	if err != nil {
		return nil, InvalidPath(path, err)
	}

	repo, err := path.Repository()
	if err != nil {
		return openOwner(ctx, gh, owner)
	}

	release, err := path.Release()
	if err != nil {
		// TODO: Content
		return openRepo(ctx, gh, owner, repo)
	}

	asset, err := path.Asset()
	if err != nil {
		return openRelease(ctx, gh, owner, repo, release)
	}

	return openAsset(ctx, gh, owner, repo, release, asset)
}

func openOwner(ctx context.Context, gh *github.Client, owner string) (afero.File, error) {
	panic("TODO")
}

func openRepo(ctx context.Context, gh *github.Client, owner, repo string) (afero.File, error) {
	panic("TODO")
}

func openRelease(ctx context.Context, gh *github.Client, owner, repo, release string) (afero.File, error) {
	panic("TODO")
}

func openAsset(ctx context.Context, gh *github.Client, owner, repo, release, asset string) (afero.File, error) {
	panic("TODO")
}

func InvalidPath(path Path, err error) error {
	return fmt.Errorf("invalid path %s: %w", path, err)
}
