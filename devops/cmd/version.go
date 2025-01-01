package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	util "github.com/unmango/go/cmd"
	"github.com/unmango/go/devops/version"
	"github.com/unmango/go/devops/work"
)

type VersionOptions struct {
	Chdir string
}

func (o VersionOptions) Cwd(ctx context.Context) (work.Directory, error) {
	if o.Chdir != "" {
		return work.Directory(o.Chdir), nil
	} else {
		return work.Load(ctx)
	}
}

func NewVersion() *cobra.Command {
	opts := VersionOptions{}

	cmd := &cobra.Command{
		Use:   "version [name]",
		Short: "Print the version of the specified dependency",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			work, err := opts.Cwd(cmd.Context())
			if err != nil {
				util.Fail(err)
			}

			if err = os.Chdir(work.Path()); err != nil {
				util.Fail(err)
			}

			v, err := version.ReadFile(args[0])
			if err != nil {
				util.Fail(err)
			}

			fmt.Println(v)
		},
	}

	cmd.Flags().StringVarP(&opts.Chdir, "chdir", "C", "", "change to the specified directory before executing")
	_ = cmd.MarkFlagDirname("chdir")

	return cmd
}
