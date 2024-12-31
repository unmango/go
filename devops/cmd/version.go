package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/go/cmd"
	"github.com/unmango/go/devops/version"
)

func NewVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of the specified dependency",
		Run: func(_ *cobra.Command, args []string) {
			if v, err := version.ReadFile(args[0]); err != nil {
				cmd.Fail(err)
			} else {
				fmt.Println(v)
			}
		},
	}
}
