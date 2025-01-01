package version

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/unmango/go/option"
)

func Init(ctx context.Context, name string, src Source, options ...Option) (err error) {
	opts := Options{fs: afero.NewOsFs()}
	option.ApplyAll(&opts, options)

	if name == "" {
		if name, err = src.Name(ctx); err != nil {
			return fmt.Errorf("name is required")
		}
	}

	if err = opts.fs.Mkdir(DirName, os.ModePerm); !errors.Is(err, os.ErrExist) {
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
