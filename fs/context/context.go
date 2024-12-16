package context

import (
	"context"

	"github.com/spf13/afero"
)

type Context = context.Context

func NewFs(base Fs, get AccessorFunc) afero.Fs {
	return &AccessorFs{get, base}
}

func BackgroundFs(base Fs) afero.Fs {
	return NewFs(base, context.Background)
}

func TodoFs(base Fs) afero.Fs {
	return NewFs(base, context.TODO)
}
