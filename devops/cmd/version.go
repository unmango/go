package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	util "github.com/unmango/go/cmd"
	"github.com/unmango/go/devops/version"
	"github.com/unmango/go/devops/work"
)

type VersionOptions struct {
	work.ChdirOptions
}

func NewVersion() *cobra.Command {
	opts := VersionOptions{}

	cmd := &cobra.Command{
		Use:   "version [name]",
		Short: "Print the version of the specified dependency",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.Chdir(cmd.Context()); err != nil {
				util.Fail(err)
			}

			v, err := version.ReadFile(args[0])
			if err != nil {
				util.Fail(err)
			}

			fmt.Println(v)
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")

	return cmd
}
