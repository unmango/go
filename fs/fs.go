package aferox

import (
	"io"

	"github.com/docker/docker/client"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/context"
	"github.com/unmango/go/fs/docker"
	"github.com/unmango/go/fs/github"
	"github.com/unmango/go/fs/github/repository"
	"github.com/unmango/go/fs/github/repository/content"
	"github.com/unmango/go/fs/github/repository/release"
	"github.com/unmango/go/fs/github/user"
	"github.com/unmango/go/fs/writer"
)

type (
	Docker                  = docker.Fs
	GitHub                  = github.Fs
	GitHubRelease           = release.Fs
	GitHubRepository        = repository.Fs
	GitHubRepositoryContent = content.Fs
	GitHubUser              = user.Fs
	Writer                  = writer.Fs
)

func NewWriter(w io.Writer) afero.Fs {
	return writer.NewFs(w)
}

func NewDocker(client client.ContainerAPIClient, container string) context.Fs {
	return docker.NewFs(client, container)
}
