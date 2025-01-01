package work

import (
	"context"

	"github.com/spf13/cobra"
)

type ChdirOptions struct {
	Chdir string
}

func (o ChdirOptions) Cwd(ctx context.Context) (Directory, error) {
	if o.Chdir != "" {
		return Directory(o.Chdir), nil
	} else {
		return Load(ctx)
	}
}

func ChdirFlag(cmd *cobra.Command, opts *ChdirOptions, value string) error {
	cmd.Flags().StringVarP(&opts.Chdir, "chdir", "C", value,
		"change to the specified directory before executing",
	)
	return cmd.MarkFlagDirname("chdir")
}
