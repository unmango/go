package internal

import (
	"errors"
	"fmt"
	"path"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/goware/urlx"
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

type Parser interface {
	Parse(string) (Path, error)
}

type OwnerPath struct {
	Owner string
}

func (p OwnerPath) Parse(path string) (Path, error) {
	return Parse(p.Owner, path)
}

func (p OwnerPath) String() string {
	return fmt.Sprintf("https://github.com/%s", p.Owner)
}

type RepositoryPath struct {
	OwnerPath
	Repository string
}

func (p RepositoryPath) Parse(path string) (Path, error) {
	if strings.HasPrefix(path, "releases/tag") {
		return Parse(p.Owner, p.Repository, path)
	} else {
		return Parse(p.Owner, p.Repository, "releases", "tag", path)
	}
}

func (p RepositoryPath) String() string {
	return fmt.Sprintf("%s/%s", p.OwnerPath, p.Repository)
}

type ReleasePath struct {
	RepositoryPath
	Release string
}

func (p ReleasePath) Parse(path string) (Path, error) {
	return Parse(p.Owner, p.Repository, "releases", "tag", p.Release, path)
}

func (p ReleasePath) String() string {
	return fmt.Sprintf("%s/releases/tag/%s", p.RepositoryPath, p.Release)
}

type AssetPath struct {
	ReleasePath
	Asset string
}

func (p AssetPath) Parse(path string) (Path, error) {
	return Parse(p.Owner, p.Repository, p.Release, p.Asset, path)
}

func (p AssetPath) String() string {
	return fmt.Sprintf("%s/download/%s", p.ReleasePath, p.Asset)
}

func NewOwnerPath(owner string) OwnerPath {
	return OwnerPath{Owner: owner}
}

func NewRepositoryPath(owner, repo string) RepositoryPath {
	return RepositoryPath{
		OwnerPath:  NewOwnerPath(owner),
		Repository: repo,
	}
}

func NewReleasePath(owner, repo, release string) ReleasePath {
	return ReleasePath{
		RepositoryPath: NewRepositoryPath(owner, repo),
		Release:        release,
	}
}

func NewAssetPath(owner, repo, release, asset string) AssetPath {
	return AssetPath{
		ReleasePath: NewReleasePath(owner, repo, release),
		Asset:       asset,
	}
}

type ghpath []string

func (g ghpath) String() string {
	return path.Join(g...)
}

// Asset implements Path.
func (g ghpath) Asset() (string, error) {
	if _, err := g.Release(); err != nil {
		return "", errors.New("not a release")
	}

	if a, err := g.index(5, "asset"); err != nil {
		return "", err
	} else if a == "download" {
		return g.index(6, "asset")
	} else {
		return a, nil
	}
}

// Branch implements Path.
func (g ghpath) Branch() (string, error) {
	if g.has(2, "tree") {
		return g.index(3, "branch")
	}

	if g.has(2, "refs") {
		return g.index(4, "branch")
	}

	return "", errors.New("not a branch")
}

// Content implements Path.
func (g ghpath) Content() []string {
	if g.has(2, "tree") {
		return g[4:]
	}

	if g.has(2, "refs") {
		return g[5:]
	}

	return []string{}
}

// Release implements Path.
func (g ghpath) Release() (string, error) {
	// This will change when I decide to support content
	if len(g) == 3 {
		return g[2], nil
	}

	if !g.has(2, "releases") {
		return "", errors.New("no release")
	}

	if g.has(3, "tag") || g.has(3, "download") {
		return g.index(4, "release")
	}

	return "", errors.New("no release")
}

// Owner implements Path.
func (g ghpath) Owner() (string, error) {
	return g.index(0, "owner")
}

// Repository implements Path.
func (g ghpath) Repository() (string, error) {
	return g.index(1, "repository")
}

func (g ghpath) has(i int, name string) bool {
	part, err := g.index(i, name)
	return err == nil && part == name
}

func (g ghpath) index(i int, name string) (string, error) {
	if len(g) <= i {
		return "", fmt.Errorf("no %s", name)
	} else {
		return g[i], nil
	}
}

func ParseUrl(rawURL string) (Path, error) {
	url, err := urlx.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(url.Path, "/")
	return Parse(parts...)
}

func Parse(parts ...string) (Path, error) {
	if len(parts) == 0 {
		return nil, errors.New("empty path")
	}

	path := []string{}
	for _, p := range parts {
		if p == "" {
			continue
		}

		url, err := urlx.Parse(p)
		if err != nil {
			log.Errorf("err: %s, p: %s", err, p)
			return nil, err
		}

		if slices.Contains(knownHosts, url.Host) {
			continue
		}

		for _, s := range strings.Split(p, "/") {
			if s == "" {
				continue
			}

			path = append(path, s)
		}
	}

	return ghpath(path), nil
}

func ParseOwner(path Path) (owner OwnerPath, err error) {
	if owner.Owner, err = path.Owner(); err != nil {
		return owner, err
	}

	return
}

func ParseRepository(path Path) (repo RepositoryPath, err error) {
	if repo.OwnerPath, err = ParseOwner(path); err != nil {
		return
	}

	if repo.Repository, err = path.Owner(); err != nil {
		return
	}

	return
}

func ParseRelease(path Path) (release ReleasePath, err error) {
	if release.RepositoryPath, err = ParseRepository(path); err != nil {
		return
	}

	if release.Release, err = path.Release(); err != nil {
		return
	}

	return
}

func ParseAsset(path Path) (asset AssetPath, err error) {
	if asset.ReleasePath, err = ParseRelease(path); err != nil {
		return
	}

	if asset.Asset, err = path.Asset(); err != nil {
		return
	}

	return
}
