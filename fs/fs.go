package fs

import (
	"github.com/unmango/go/fs/github"
	"github.com/unmango/go/fs/github/repository"
	"github.com/unmango/go/fs/github/repository/content"
	"github.com/unmango/go/fs/github/repository/release"
	"github.com/unmango/go/fs/github/user"
)

type (
	GitHub                  = github.Fs
	GitHubUser              = user.Fs
	GitHubRepository        = repository.Fs
	GitHubRepositoryContent = content.Fs
	GitHubRelease           = release.Fs
)
