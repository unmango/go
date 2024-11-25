package internal

import (
	"errors"
	"fmt"
	"net/url"
	"path"
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

type OwnerPath struct {
	Owner string
}

func (p OwnerPath) String() string {
	return fmt.Sprintf("https://github.com/%s", p.Owner)
}

type RepositoryPath struct {
	OwnerPath
	Repository string
}

func (p RepositoryPath) String() string {
	return fmt.Sprintf("%s/%s", p.OwnerPath, p.Repository)
}

type ReleasePath struct {
	RepositoryPath
	Release string
}

func (p ReleasePath) String() string {
	return fmt.Sprintf("%s/releases/tag/%s", p.RepositoryPath, p.Release)
}

type AssetPath struct {
	ReleasePath
	Asset string
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

func Parse(parts ...string) (Path, error) {
	path := path.Join(parts...)
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

func ParseOwner(path string) (owner OwnerPath, err error) {
	if p, err := Parse(path); err != nil {
		return owner, err
	} else {
		return parseOwner(p)
	}
}

func ParseRepository(path string) (repo RepositoryPath, err error) {
	if p, err := Parse(path); err != nil {
		return repo, err
	} else {
		return parseRepository(p)
	}
}

func ParseRelease(path string) (release ReleasePath, err error) {
	if p, err := Parse(path); err != nil {
		return release, err
	} else {
		return parseRelease(p)
	}
}

func ParseAsset(path string) (asset AssetPath, err error) {
	if p, err := Parse(path); err != nil {
		return asset, err
	} else {
		return parseAsset(p)
	}
}

func parseOwner(path Path) (owner OwnerPath, err error) {
	if owner.Owner, err = path.Owner(); err != nil {
		return owner, err
	}

	return
}

func parseRepository(path Path) (repo RepositoryPath, err error) {
	if repo.OwnerPath, err = parseOwner(path); err != nil {
		return
	}

	if repo.Repository, err = path.Owner(); err != nil {
		return
	}

	return
}

func parseRelease(path Path) (release ReleasePath, err error) {
	if release.RepositoryPath, err = parseRepository(path); err != nil {
		return
	}

	if release.Release, err = path.Release(); err != nil {
		return
	}

	return
}

func parseAsset(path Path) (asset AssetPath, err error) {
	if asset.ReleasePath, err = parseRelease(path); err != nil {
		return
	}

	if asset.Asset, err = path.Asset(); err != nil {
		return
	}

	return
}
