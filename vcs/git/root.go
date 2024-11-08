package git

import (
	"context"
	"os/exec"
	"strings"

	"github.com/unmango/go/option"
)

type rootOption func(*exec.Cmd)

func Root(ctx context.Context, options ...rootOption) (string, error) {
	cmd := exec.CommandContext(ctx,
		"git", "rev-parse", "--show-toplevel",
	)
	option.ApplyAll(cmd, options)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func WithWorkingDirectory(path string) rootOption {
	return func(c *exec.Cmd) {
		c.Dir = path
	}
}
