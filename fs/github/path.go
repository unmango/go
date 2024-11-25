package github

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var knownHosts = []string{
	"github.com",
	"api.github.com",
	"raw.githubusercontent.com",
}

type Path interface {
	fmt.Stringer
	Asset() (string, error)
	Branch() (string, error)
	Content() []string
	Owner() (string, error)
	Repository() (string, error)
	Release() (string, error)
}

type ghpath struct{ *url.URL }

// Asset implements Path.
func (g *ghpath) Asset() (string, error) {
	if _, err := g.Release(); err != nil {
		return "", errors.New("not a release")
	}

	return g.index(5, "asset")
}

// Branch implements Path.
func (g *ghpath) Branch() (string, error) {
	if g.has(2, "tree") {
		return g.index(3, "branch")
	}

	if g.has(2, "refs") {
		return g.index(4, "branch")
	}

	return "", errors.New("not a branch")
}

// Content implements Path.
func (g *ghpath) Content() []string {
	if g.has(2, "tree") {
		return g.parts()[4:]
	}

	if g.has(2, "refs") {
		return g.parts()[5:]
	}

	return []string{}
}

// Release implements Path.
func (g *ghpath) Release() (string, error) {
	if !g.has(2, "releases") {
		return "", errors.New("no release")
	}

	if g.has(3, "tag") || g.has(3, "download") {
		return g.index(4, "release")
	}

	return "", errors.New("no release")
}

// Owner implements Path.
func (g *ghpath) Owner() (string, error) {
	return g.index(0, "owner")
}

// Repository implements Path.
func (g *ghpath) Repository() (string, error) {
	return g.index(1, "repository")
}

func (g *ghpath) parts() []string {
	return strings.Split(strings.TrimPrefix(g.Path, "/"), "/")
}

func (g *ghpath) has(i int, name string) bool {
	part, err := g.index(i, name)
	return err == nil && part == name
}

func (g *ghpath) index(i int, name string) (string, error) {
	if parts := g.parts(); len(parts) <= i {
		return "", fmt.Errorf("no %s", name)
	} else {
		return parts[i], nil
	}
}

func Parse(path string) (Path, error) {
	for _, host := range knownHosts {
		if strings.HasPrefix(path, host) {
			path = fmt.Sprintf("https://%s", path)
		}
	}

	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	return &ghpath{url}, nil
}
