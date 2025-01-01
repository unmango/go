package git

import (
	"context"
	"os/exec"
	"strings"
)

type (
	wdKey struct{}
)

func Root(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx,
		"git", "rev-parse", "--show-toplevel",
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

func WithWorkingDirectory(parent context.Context, path string) context.Context {
	return context.WithValue(parent, wdKey{}, path)
}
