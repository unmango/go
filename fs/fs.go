package fs

import (
	"io"

	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github"
	"github.com/unmango/go/fs/github/repository"
	"github.com/unmango/go/fs/github/repository/content"
	"github.com/unmango/go/fs/github/repository/release"
	"github.com/unmango/go/fs/github/user"
	"github.com/unmango/go/fs/writer"
)

type (
	GitHub                  = github.Fs
	GitHubUser              = user.Fs
	GitHubRepository        = repository.Fs
	GitHubRepositoryContent = content.Fs
	GitHubRelease           = release.Fs
	Writer                  = writer.Fs
)

func NewWriter(w io.Writer) afero.Fs {
	return writer.New(w)
}
