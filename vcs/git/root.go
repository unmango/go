package git

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

type (
	wdKey  struct{}
	gitKey struct{}
)

func Root(ctx context.Context) (string, error) {
	git, err := gitPath(ctx)
	if err != nil {
		return "", err
	}

	cmd := exec.CommandContext(ctx,
		git, "rev-parse", "--show-toplevel",
	)
	if wd, ok := ctx.Value(wdKey{}).(string); ok && wd != "" {
		cmd.Dir = wd
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func WithGitPath(parent context.Context, path string) context.Context {
	return context.WithValue(parent, gitKey{}, path)
}

func WithWorkingDirectory(parent context.Context, path string) context.Context {
	return context.WithValue(parent, wdKey{}, path)
}

func gitPath(ctx context.Context) (string, error) {
	if git := ctx.Value(gitKey{}); git != nil {
		return git.(string), nil
	}

	if git := os.Getenv("GIT_PATH"); git != "" {
		return git, nil
	}

	return exec.LookPath("git")
}
