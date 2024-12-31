package work

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/unmango/go/vcs/git"
)

type cwdKey struct{}

type Context struct {
	root string
}

func (c *Context) Root() string {
	return c.root
}

func (c *Context) JoinPath(elem ...string) string {
	parts := append([]string{c.root}, elem...)
	return filepath.Join(parts...)
}

func Git(ctx context.Context) (*Context, error) {
	if p, err := git.Root(ctx); err != nil {
		return nil, err
	} else {
		return &Context{root: p}, nil
	}
}

func Cwd() (*Context, error) {
	if p, err := os.Getwd(); err != nil {
		return nil, err
	} else {
		return &Context{root: p}, nil
	}
}

func Load(ctx context.Context) (work *Context, err error) {
	if work, err = Git(ctx); err == nil {
		return
	} else {
		log.Debugf("loading git context: %s", err)
	}

	if work, err = Cwd(); err == nil {
		return
	} else {
		log.Debugf("loading cwd context: %s", err)
	}

	return nil, fmt.Errorf("failed to load current work context")
}
