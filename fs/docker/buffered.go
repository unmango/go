package docker

import (
	"github.com/docker/docker/client"
	"github.com/spf13/afero"
)

// TODO: For some reason I thought Sync() existed on [afero.Fs] but
// it's actually on [afero.File], I'll have to re-think how I want
// this to work

type BufferedFs struct {
	afero.Fs
	layer afero.Fs
}

func NewBuffered(client client.ContainerAPIClient, container string) afero.Fs {
	layer := afero.NewMemMapFs()
	docker := afero.NewCopyOnWriteFs(
		NewFs(client, container),
		layer,
	)

	return &BufferedFs{
		Fs:    docker,
		layer: layer,
	}
}
