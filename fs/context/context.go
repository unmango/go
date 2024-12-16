package context

import (
	"context"

	"github.com/spf13/afero"
	aferox "github.com/unmango/go/fs"
)

type Context = context.Context

var (
	Background        = context.Background
	TODO              = context.TODO
	WithCancel        = context.WithCancel
	WithCancelCause   = context.WithCancelCause
	WithDeadline      = context.WithDeadline
	WithDeadlineCause = context.WithDeadlineCause
	WithValue         = context.WithValue
	WithTimeout       = context.WithTimeout
	WithTimeoutCause  = context.WithTimeoutCause
	WithoutCancel     = context.WithoutCancel
)

type union struct {
	afero.Fs
	AdapterFs
}

func NewFs(base Fs, accessor aferox.ContextAccessor) afero.Fs {
	return &AccessorFs{accessor, base}
}

func BackgroundFs(base Fs) afero.Fs {
	return NewFs(base, AccessorFunc(Background))
}

func TodoFs(base Fs) afero.Fs {
	return NewFs(base, AccessorFunc(TODO))
}

func Discard(fs afero.Fs) AferoFs {
	return &DiscardFs{fs}
}

func Adapt(fs Fs, accessor aferox.ContextAccessor) AferoFs {
	return &union{
		Fs:        NewFs(fs, accessor),
		AdapterFs: AdapterFs{fs},
	}
}
