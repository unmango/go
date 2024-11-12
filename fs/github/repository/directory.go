package repository

import (
	"context"
	"fmt"

	"github.com/google/go-github/v66/github"
)

func Readdir(ctx context.Context, gh *github.Client, user string, count int) ([]*FileInfo, error) {
	repos, _, err := gh.Repositories.ListByUser(ctx, user, nil)
	if err != nil {
		return nil, fmt.Errorf("user %s readdir: %w", user, err)
	}

	length := min(count, len(repos))
	results := make([]*FileInfo, length)

	for i := 0; i < length; i++ {
		results[i] = &FileInfo{
			client: gh,
			repo:   repos[i],
		}
	}

	return results, nil
}
