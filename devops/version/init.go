package version

import (
	"context"
	"os"

	"github.com/spf13/afero"
	"github.com/unmango/go/option"
)

const DirName = ".versions"

type InitOptions struct {
	fs afero.Fs
}

type InitOption func(*InitOptions)

func Init(ctx context.Context, name string, src Source, options ...InitOption) (err error) {
	opts := InitOptions{fs: afero.NewOsFs()}
	option.ApplyAll(&opts, options)

	if name == "" {
		if name, err = src.Name(ctx); err != nil {
			return
		}
	}

	if err = opts.fs.Mkdir(DirName, os.ModePerm); err != nil {
		return
	}

	var version string
	if version, err = src.Latest(ctx); err != nil {
		return
	}

	return afero.WriteFile(opts.fs,
		RelPath(name),
		[]byte(version+"\n"),
		os.ModePerm,
	)
}

func WithInitFs(fs afero.Fs) InitOption {
	return func(o *InitOptions) {
		o.fs = fs
	}
}
