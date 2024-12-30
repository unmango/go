package cmd

import "github.com/spf13/cobra"

func NewVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of the specified dependency",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}
