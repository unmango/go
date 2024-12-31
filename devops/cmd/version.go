package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	util "github.com/unmango/go/cmd"
	"github.com/unmango/go/devops/version"
	"github.com/unmango/go/devops/work"
)

func NewVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of the specified dependency",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			work, err := work.Load(cmd.Context())
			if err != nil {
				util.Fail(err)
			}

			v, err := version.ReadFile(args[0],
				version.WithWorkspace(work),
			)
			if err != nil {
				util.Fail(err)
			}

			fmt.Println(v)
		},
	}
}
